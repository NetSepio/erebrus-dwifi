package wifiAlgorithem

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"net"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/NetSepio/erebrus-dwifi/connections/database/wifiDbOperations"
	"github.com/fatih/color"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

type deviceInfo struct {
	macAddress         net.HardwareAddr
	ipAddress          net.IP
	connectedAt        time.Time
	totalConnectedTime time.Duration
	connected          bool
	lastChecked        time.Time
	defaultGateway     string
	manufacturer       string
	interfaceName      string
	connectionType     string
}

var devices = make(map[string]*deviceInfo)
var devicesLock sync.Mutex

func RunWifi() {
	color.NoColor = false // Enable color output even if output is not a terminal

	ifaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	for _, iface := range ifaces {
		wg.Add(1)
		go func(iface net.Interface) {
			defer wg.Done()
			if err := scan(&iface); err != nil {
				log.Printf("interface %v: %v", iface.Name, err)
			}
		}(iface)
	}
	wg.Wait()

	printTotalConnectedTime()
}

func scan(iface *net.Interface) error {
	var addr *net.IPNet
	if addrs, err := iface.Addrs(); err != nil {
		return err
	} else {
		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok {
				if ip4 := ipnet.IP.To4(); ip4 != nil {
					addr = &net.IPNet{
						IP:   ip4,
						Mask: ipnet.Mask[len(ipnet.Mask)-4:],
					}
					break
				}
			}
		}
	}

	if addr == nil {
		return errors.New("no good IP network found")
	} else if addr.IP[0] == 127 {
		return errors.New("skipping localhost")
	} else if addr.Mask[0] != 0xff || addr.Mask[1] != 0xff {
		return errors.New("mask means network is too large")
	}
	log.Printf("Using network range %v for interface %v", addr, iface.Name)

	handle, err := pcap.OpenLive(iface.Name, 65536, true, pcap.BlockForever)
	if err != nil {
		return err
	}
	defer handle.Close()

	stop := make(chan struct{})
	go readARP(handle, iface, stop)
	defer close(stop)

	go func() {
		for {
			time.Sleep(15 * time.Second)
			checkDisconnectedDevices()
		}
	}()

	for {
		if err := writeARP(handle, iface, addr); err != nil {
			log.Printf("error writing packets on %v: %v", iface.Name, err)
			return err
		}
		time.Sleep(10 * time.Second)
	}
}

func readARP(handle *pcap.Handle, iface *net.Interface, stop chan struct{}) {
	src := gopacket.NewPacketSource(handle, layers.LayerTypeEthernet)
	in := src.Packets()
	for {
		select {
		case <-stop:
			return
		case packet := <-in:
			arpLayer := packet.Layer(layers.LayerTypeARP)
			if arpLayer == nil {
				continue
			}
			arp := arpLayer.(*layers.ARP)
			if arp.Operation != layers.ARPReply || bytes.Equal([]byte(iface.HardwareAddr), arp.SourceHwAddress) {
				continue
			}

			mac := net.HardwareAddr(arp.SourceHwAddress).String()
			devicesLock.Lock()
			device, exists := devices[mac]
			if !exists {
				manufacturer := getManufacturer(mac)
				defaultGateway := getDefaultGateway()
				interfaceName := iface.Name
				connectionType := getConnectionType(interfaceName)

				devices[mac] = &deviceInfo{
					macAddress:         net.HardwareAddr(arp.SourceHwAddress),
					ipAddress:          net.IP(arp.SourceProtAddress),
					connectedAt:        time.Now(),
					totalConnectedTime: 0,
					connected:          true,
					lastChecked:        time.Now(),
					defaultGateway:     defaultGateway,
					manufacturer:       manufacturer,
					interfaceName:      interfaceName,
					connectionType:     connectionType,
				}
				logDeviceInfo("Device connected", devices[mac])
			} else if !device.connected {
				device.ipAddress = net.IP(arp.SourceProtAddress)
				device.connectedAt = time.Now()
				device.connected = true
				device.lastChecked = time.Now()
				logDeviceInfo("Device reconnected", device)
			}
			devicesLock.Unlock()
		}
	}
}

func writeARP(handle *pcap.Handle, iface *net.Interface, addr *net.IPNet) error {
	eth := layers.Ethernet{
		SrcMAC:       iface.HardwareAddr,
		DstMAC:       net.HardwareAddr{0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		EthernetType: layers.EthernetTypeARP,
	}
	arp := layers.ARP{
		AddrType:          layers.LinkTypeEthernet,
		Protocol:          layers.EthernetTypeIPv4,
		HwAddressSize:     6,
		ProtAddressSize:   4,
		Operation:         layers.ARPRequest,
		SourceHwAddress:   []byte(iface.HardwareAddr),
		SourceProtAddress: []byte(addr.IP),
		DstHwAddress:      []byte{0, 0, 0, 0, 0, 0},
	}
	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}
	for _, ip := range ips(addr) {
		arp.DstProtAddress = []byte(ip)
		gopacket.SerializeLayers(buf, opts, &eth, &arp)
		if err := handle.WritePacketData(buf.Bytes()); err != nil {
			return err
		}
	}
	return nil
}

func ips(n *net.IPNet) (out []net.IP) {
	num := binary.BigEndian.Uint32([]byte(n.IP))
	mask := binary.BigEndian.Uint32([]byte(n.Mask))
	network := num & mask
	broadcast := network | ^mask
	for network++; network < broadcast; network++ {
		var buf [4]byte
		binary.BigEndian.PutUint32(buf[:], network)
		out = append(out, net.IP(buf[:]))
	}
	return
}

func checkDisconnectedDevices() {
	devicesLock.Lock()
	defer devicesLock.Unlock()
	for _, device := range devices {
		if device.connected && time.Since(device.lastChecked) > 15*time.Second {
			device.totalConnectedTime += time.Since(device.connectedAt)
			device.connected = false
			logDeviceInfo("Device disconnected", device)
		}
	}
}

func printTotalConnectedTime() {
	devicesLock.Lock()
	defer devicesLock.Unlock()
	for _, device := range devices {
		if device.totalConnectedTime > 0 {
			logDeviceInfo("Total connected time", device)
		}
	}
}

func getDefaultGateway() string {
	out, err := exec.Command("ip", "route", "show", "default").Output()
	if err != nil {
		return "Unknown"
	}
	parts := strings.Fields(string(out))
	if len(parts) > 2 {
		return parts[2]
	}
	return "Unknown"
}

func getManufacturer(mac string) string {
	out, err := exec.Command("sudo", "arp-scan", "--localnet").Output()
	if err != nil {
		return "Unknown Manufacturer"
	}

	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if strings.Contains(line, mac) {
			fields := strings.Fields(line)
			if len(fields) > 2 {
				return strings.Join(fields[2:], " ")
			}
		}
	}
	return "Unknown Manufacturer"
}

func logDeviceInfo(event string, device *deviceInfo) {

	// ADD THIS

	data := wifiDbOperations.DeviceInfo{
		MacAddress:         device.macAddress.String(),
		Manufacturer:       device.manufacturer,
		TotalConnectedTime: device.totalConnectedTime.String(),
	}

	WifiCollector(data)

	fmt.Println("----------------------------------------")
	fmt.Printf("%s : %s\n", color.HiGreenString(event), color.HiBlueString("Device"))

	fmt.Printf("  MAC: %s\n", color.HiMagentaString(device.macAddress.String()))
	fmt.Printf("  IP: %s\n", color.HiMagentaString(device.ipAddress.String()))
	fmt.Printf("  Manufacturer: %s\n", color.HiRedString(device.manufacturer))
	fmt.Printf("  Gateway: %s\n", color.HiYellowString(device.defaultGateway))
	fmt.Printf("  Interface: %s\n", color.HiYellowString(device.interfaceName))
	fmt.Printf("  Host SSID: %s\n", color.HiYellowString(device.connectionType))
	fmt.Printf("  Total connected time: %v\n", color.HiCyanString(device.totalConnectedTime.String()))
}

func getConnectionType(interfaceName string) string {
	out, err := exec.Command("nmcli", "dev", "show", interfaceName).Output()
	if err != nil {
		return "Unknown"
	}
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if strings.Contains(line, "GENERAL.CONNECTION") {
			fields := strings.Fields(line)
			if len(fields) > 1 {
				return fields[1]
			}
		}
	}
	return "Unknown"
}

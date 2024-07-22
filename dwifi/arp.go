package dwifi

import (
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/fatih/color"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func Scan(iface *net.Interface) error {
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
	defaultGateway, _ := getPublicIP()
	fmt.Printf(" Default Gateway: %s\n", color.HiYellowString(defaultGateway))

	fmt.Printf(" Using network range %v for interface %v\n", addr, iface.Name)

	handle, err := pcap.OpenLive(iface.Name, 65536, true, pcap.BlockForever)
	if err != nil {
		return err
	}
	defer handle.Close()

	stop := make(chan struct{})
	go readARP(handle, iface, stop, defaultGateway)
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

// func readARP(handle *pcap.Handle, iface *net.Interface, stop chan struct{}) {
// 	src := gopacket.NewPacketSource(handle, layers.LayerTypeEthernet)
// 	in := src.Packets()
// 	for {
// 		select {
// 		case <-stop:
// 			return
// 		case packet := <-in:
// 			arpLayer := packet.Layer(layers.LayerTypeARP)
// 			if arpLayer == nil {
// 				continue
// 			}
// 			arp := arpLayer.(*layers.ARP)
// 			if arp.Operation != layers.ARPReply || bytes.Equal([]byte(iface.HardwareAddr), arp.SourceHwAddress) {
// 				continue
// 			}

// 			mac := net.HardwareAddr(arp.SourceHwAddress).String()
// 			ip := net.IP(arp.SourceProtAddress).String()
// 			devicesLock.Lock()
// 			device, exists := devices[mac]
// 			if !exists {
// 				manufacturer := getManufacturer(mac)
// 				defaultGateway := getDefaultGateway()
// 				interfaceName := iface.Name
// 				hostSSID := getHostSSID(interfaceName)

// 				devices[mac] = &deviceInfo{
// 					MACAddress:         mac,
// 					IPAddress:          ip,
// 					ConnectedAt:        time.Now(),
// 					TotalConnectedTime: 0,
// 					Connected:          true,
// 					LastChecked:        time.Now(),
// 					DefaultGateway:     defaultGateway,
// 					Manufacturer:       manufacturer,
// 					InterfaceName:      interfaceName,
// 					HostSSID:           hostSSID,
// 				}
// 				logDeviceInfo("Device connected", devices[mac])
// 			} else if !device.Connected {
// 				device.IPAddress = ip
// 				device.ConnectedAt = time.Now()
// 				device.Connected = true
// 				device.LastChecked = time.Now()
// 				// logDeviceInfo("Device reconnected", device)
// 			}
// 			devicesLock.Unlock()
// 		}
// 	}
// }

func readARP(handle *pcap.Handle, iface *net.Interface, stop chan struct{}, defaultGateway string) {
	src := gopacket.NewPacketSource(handle, layers.LayerTypeEthernet)
	in := src.Packets()

	// Add the device running the code
	ownIP := getOwnIP(iface)
	if ownIP != "" {
		addDevice(iface.HardwareAddr.String(), ownIP, defaultGateway, iface.Name)
	}

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
			if arp.Operation != layers.ARPReply {
				continue
			}

			mac := net.HardwareAddr(arp.SourceHwAddress).String()
			ip := net.IP(arp.SourceProtAddress).String()
			addDevice(mac, ip, defaultGateway, iface.Name)
		}
	}
}

func addDevice(mac, ip, defaultGateway, interfaceName string) {
	devicesLock.Lock()
	defer devicesLock.Unlock()

	device, exists := devices[mac]
	if !exists {
		manufacturer := getManufacturer(mac)
		hostSSID := getHostSSID(interfaceName)

		devices[mac] = &deviceInfo{
			MACAddress:         mac,
			IPAddress:          ip,
			ConnectedAt:        time.Now(),
			TotalConnectedTime: 0,
			Connected:          true,
			LastChecked:        time.Now(),
			DefaultGateway:     defaultGateway,
			Manufacturer:       manufacturer,
			InterfaceName:      interfaceName,
			HostSSID:           hostSSID,
		}

		logDeviceInfo("Device detected", devices[mac])
	} else if !device.Connected {
		device.IPAddress = ip
		device.ConnectedAt = time.Now()
		device.Connected = true
		device.LastChecked = time.Now()
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

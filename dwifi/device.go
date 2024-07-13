package dwifi

import (
	"net"
	"sync"
	"time"
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
	HostSSID           string
}

var devices = make(map[string]*deviceInfo)
var devicesLock sync.Mutex

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

func PrintTotalConnectedTime() {
	devicesLock.Lock()
	defer devicesLock.Unlock()
	for _, device := range devices {
		if device.totalConnectedTime > 0 {
			logDeviceInfo("Total connected time", device)
		}
	}
}

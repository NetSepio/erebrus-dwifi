// package dwifi

// import (
// 	"encoding/json"
// 	"log"
// 	"net"
// 	"sync"
// 	"time"

// 	"gorm.io/gorm"
// )

// type deviceInfo struct {
// 	MACAddress         net.HardwareAddr `json:"macAddress"`
// 	IPAddress          net.IP           `json:"ipAddress"`
// 	ConnectedAt        time.Time        `json:"connectedAt"`
// 	TotalConnectedTime time.Duration    `json:"totalConnectedTime"`
// 	Connected          bool             `json:"connected"`
// 	LastChecked        time.Time        `json:"lastChecked"`
// 	DefaultGateway     string           `json:"defaultGateway"`
// 	Manufacturer       string           `json:"manufacturer"`
// 	InterfaceName      string           `json:"interfaceName"`
// 	HostSSID           string           `json:"hostSSID"`
// }

// var devices = make(map[string]*deviceInfo)
// var devicesLock sync.Mutex
// var db *gorm.DB

// func InitDevicesDatabase(database *gorm.DB) {
// 	db = database
// }

// func checkDisconnectedDevices() {
// 	devicesLock.Lock()
// 	defer devicesLock.Unlock()
// 	for _, device := range devices {
// 		if device.Connected && time.Since(device.LastChecked) > 15*time.Second {
// 			device.TotalConnectedTime += time.Since(device.ConnectedAt)
// 			device.Connected = false
// 			logDeviceInfo("Device disconnected", device)
// 			updateDatabase(device)
// 		}
// 	}
// }

// func PrintTotalConnectedTime() {
// 	devicesLock.Lock()
// 	defer devicesLock.Unlock()
// 	for _, device := range devices {
// 		if device.TotalConnectedTime > 0 {
// 			logDeviceInfo("Total connected time", device)
// 			updateDatabase(device)
// 		}
// 	}
// }

// func updateDatabase(device *deviceInfo) {
// 	defaultGateway := device.DefaultGateway
// 	devicesByGateway := getDevicesByGateway(defaultGateway)
// 	jsonData, err := json.Marshal(devicesByGateway)
// 	if err != nil {
// 		log.Printf("Failed to marshal device data: %v", err)
// 		return
// 	}

// 	nodeData := NodeDwifi{
// 		Gateway: defaultGateway,
// 		Status:  string(jsonData),
// 	}

// 	if err := SaveNodeData(db, &nodeData); err != nil {
// 		log.Printf("Failed to save node data: %v", err)
// 	}
// }

// func getDevicesByGateway(defaultGateway string) []*deviceInfo {
// 	var result []*deviceInfo
// 	for _, device := range devices {
// 		if device.DefaultGateway == defaultGateway {
// 			result = append(result,
// 			device.MACAddress: net.HardwareAddr.String(),
// 			device.IPAddress : net.IP.String(),

// 			)
// 		}
// 	}
// 	return result
// }

package dwifi

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"gorm.io/gorm"
)

type deviceInfo struct {
	MACAddress         string    `json:"macAddress"`
	IPAddress          string    `json:"ipAddress"`
	ConnectedAt        time.Time `json:"connectedAt"`
	TotalConnectedTime float64   `json:"totalConnectedTime"`
	Connected          bool      `json:"connected"`
	LastChecked        time.Time `json:"lastChecked"`
	DefaultGateway     string    `json:"defaultGateway"`
	Manufacturer       string    `json:"manufacturer"`
	InterfaceName      string    `json:"interfaceName"`
	HostSSID           string    `json:"hostSSID"`
}

var devices = make(map[string]*deviceInfo)
var devicesLock sync.Mutex
var db *gorm.DB
var walletAddress string
var pricePerMin float64

func InitDevicesDatabase(database *gorm.DB, wallet string, price float64) {
	db = database
	walletAddress = wallet
	pricePerMin = price
}

func checkDisconnectedDevices() {
	devicesLock.Lock()
	defer devicesLock.Unlock()
	for _, device := range devices {
		if device.Connected && time.Since(device.LastChecked) > 15*time.Second {
			device.TotalConnectedTime += float64(time.Since(device.ConnectedAt))
			device.Connected = false
			logDeviceInfo("Device disconnected", device)
			updateDatabase(device)
		}
	}
}

func PrintTotalConnectedTime() {
	devicesLock.Lock()
	defer devicesLock.Unlock()
	for _, device := range devices {
		if device.TotalConnectedTime > 0 {
			logDeviceInfo(fmt.Sprintf("Total connected time: %.2f minutes", device.TotalConnectedTime), device)
			updateDatabase(device)
		}
	}
}

func updateDatabase(device *deviceInfo) {
	publicIP, err := getPublicIP()
	if err != nil {
		log.Printf("Failed to get public IP: %v", err)
		publicIP = "Unknown"
	}

	location, err := getLocation()
	if err != nil {
		log.Printf("Failed to get location: %v", err)
		location = "Unknown"
	}

	defaultGateway := device.DefaultGateway
	devicesByGateway := getDevicesByGateway(defaultGateway)

	displayDevices := make([]*deviceInfo, len(devicesByGateway))
	for i, d := range devicesByGateway {
		displayDevice := *d // Create a copy
		displayDevice.TotalConnectedTime = float64(d.TotalConnectedTime) / float64(time.Minute)
		displayDevices[i] = &displayDevice
	}
	jsonData, err := json.Marshal(displayDevices)
	if err != nil {
		log.Printf("Failed to marshal device data: %v", err)
		return
	}

	_, password, err := getWiFiPassword()
	if err != nil {
		log.Printf("Failed to get WiFi password: %v", err)
		password = "Unknown"
	}

	nodeData := NodeDwifi{
		Gateway:       publicIP,
		Password:      password,
		WalletAddress: walletAddress,
		Location:      location,
		PricePerMin:   pricePerMin,
		Status:        string(jsonData),
	}

	if err := SaveNodeData(db, &nodeData); err != nil {
		log.Printf("Failed to save node data: %v", err)
	}
}

func getDevicesByGateway(defaultGateway string) []*deviceInfo {
	var result []*deviceInfo
	for _, device := range devices {
		if device.DefaultGateway == defaultGateway {
			result = append(result, &deviceInfo{
				MACAddress:         device.MACAddress,
				IPAddress:          device.IPAddress,
				ConnectedAt:        device.ConnectedAt,
				TotalConnectedTime: device.TotalConnectedTime,
				Connected:          device.Connected,
				LastChecked:        device.LastChecked,
				DefaultGateway:     device.DefaultGateway,
				Manufacturer:       device.Manufacturer,
				InterfaceName:      device.InterfaceName,
				HostSSID:           device.HostSSID,
			})
		}
	}
	return result
}

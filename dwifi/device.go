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
	MACAddress         string        `json:"macAddress"`
	IPAddress          string        `json:"ipAddress"`
	ConnectedAt        time.Time     `json:"connectedAt"`
	TotalConnectedTime time.Duration `json:"totalConnectedTime"`
	Connected          bool          `json:"connected"`
	LastChecked        time.Time     `json:"lastChecked"`
	DefaultGateway     string        `json:"defaultGateway"`
	Manufacturer       string        `json:"manufacturer"`
	InterfaceName      string        `json:"interfaceName"`
	HostSSID           string        `json:"hostSSID"`
}

var devices = make(map[string]*deviceInfo)
var devicesLock sync.Mutex
var db *gorm.DB
var walletAddress string
var pricePerMin string

func InitDevicesDatabase(database *gorm.DB, wallet string, price string) {
	db = database
	walletAddress = wallet
	pricePerMin = price
}

func checkDisconnectedDevices() {
	devicesLock.Lock()
	defer devicesLock.Unlock()
	for _, device := range devices {
		if device.Connected && time.Since(device.LastChecked) > 15*time.Second {
			device.TotalConnectedTime += time.Since(device.ConnectedAt)
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
			logDeviceInfo(fmt.Sprintf("Total connected time: %s", device.TotalConnectedTime), device)
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
		Gateway:        publicIP,
		Password:       password,
		Wallet_Address: walletAddress,
		Location:       location,
		Price_per_min:  pricePerMin,
		Status:         string(jsonData),
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

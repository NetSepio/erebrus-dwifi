// package wifiDbOperations

// import (
// 	"time"

// 	"github.com/google/uuid"
// )

// type DeviceInfo struct {
// 	Id                 uuid.UUID
// 	PeerId             string
// 	MacAddress         string
// 	IpAddress          string
// 	ConnectedAt        time.Time
// 	TotalConnectedTime string
// 	Connected          bool
// 	LastChecked        time.Time
// 	DefaultGateway     string
// 	Manufacturer       string
// 	InterfaceName      string
// 	HostSSID           string
// }

package wifiDbOperations

import (
	"time"

	"github.com/google/uuid"
)

type DeviceInfo struct {
	Id     uuid.UUID `gorm:"type:uuid"`
	Status string    `gorm:"type:jsonb"` // Store device info as JSON
}

// Additional struct to hold device information
type DeviceStatus struct {
	MacAddress         string    `json:"mac_address"`
	IpAddress          string    `json:"ip_address"`
	ConnectedAt        time.Time `json:"connected_at"`
	TotalConnectedTime string    `json:"total_connected_time"`
	Connected          bool      `json:"connected"`
	LastChecked        time.Time `json:"last_checked"`
	DefaultGateway     string    `json:"default_gateway"`
	Manufacturer       string    `json:"manufacturer"`
	InterfaceName      string    `json:"interface_name"`
	HostSSID           string    `json:"host_ssid"`
}

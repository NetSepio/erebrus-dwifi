package wifiDbOperations

import "github.com/google/uuid"

type DeviceInfo struct {
	Id                 uuid.UUID
	PeerId             string
	MacAddress         string
	IpAddress          string
	ConnectedAt        string
	TotalConnectedTime string
	Connected          bool
	LastChecked        string
	DefaultGateway     string
	Manufacturer       string
	InterfaceName      string
	ConnectionType     string
}

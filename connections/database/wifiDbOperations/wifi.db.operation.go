package wifiDbOperations

import (
	"github.com/NetSepio/erebrus-dwifi/connections/database"
	"github.com/google/uuid"
)

func CreateDevice(device *DeviceInfo) error {
	db := database.Connect()
	device.Id = uuid.New() // Ensure the ID is set
	result := db.Create(device)
	return result.Error
}

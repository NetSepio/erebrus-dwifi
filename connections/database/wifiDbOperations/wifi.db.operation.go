// // package wifiDbOperations

// // import (
// // 	"log"

// // 	"github.com/NetSepio/erebrus-dwifi/connections/database"
// // 	"github.com/google/uuid"
// // )

// // func CreateDevice(device *DeviceInfo) error {
// // 	db := database.Connect()
// // 	device.Id = uuid.New() // Ensure the ID is set
// // 	err := db.AutoMigrate(&DeviceInfo{})
// // 	if err != nil {
// // 		log.Fatalf("failed to migrate tables: %v", err)
// // 	}
// // 	result := db.Create(device)
// // 	return result.Error
// // }

// package wifiDbOperations

// import (
// 	"encoding/json"

// 	"github.com/NetSepio/erebrus-dwifi/connections/database"
// 	"github.com/google/uuid"
// )

// func CreateDevice(device *DeviceStatus) error {
// 	db := database.Connect()

// 	statusJson, err := json.Marshal(device)
// 	if err != nil {
// 		return err
// 	}
// 	db.AutoMigrate(&DeviceInfo{})

// 	deviceInfo := &DeviceInfo{
// 		Id:     uuid.New(), // Ensure the ID is set
// 		Status: string(statusJson),
// 	}
// 	result := db.Create(deviceInfo)
// 	return result.Error
// }

package wifiDbOperations

import (
	"encoding/json"

	"github.com/NetSepio/erebrus-dwifi/connections/database"
	"github.com/google/uuid"
)

func CreateDevice(device *DeviceStatus) error {
	db := database.Connect()

	// Ensure the table schema is up-to-date
	if err := db.AutoMigrate(&DeviceInfo{}); err != nil {
		return err
	}

	// Marshal the device status to JSON
	statusJson, err := json.Marshal(device)
	if err != nil {
		return err
	}

	deviceInfo := &DeviceInfo{
		Id:     uuid.New(), // Ensure the ID is set
		Status: string(statusJson),
	}

	// Save the device info to the database
	result := db.Create(deviceInfo)
	return result.Error
}

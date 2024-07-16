package dwifi

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type NodeDwifi struct {
	ID        uint   `gorm:"primaryKey"`
	Gateway   string `gorm:"index"`
	Status    string `gorm:"type:jsonb"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func InitDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&NodeDwifi{}); err != nil {
		return nil, err
	}

	return db, nil
}

func SaveNodeData(db *gorm.DB, node *NodeDwifi) error {
	var existing NodeDwifi
	if err := db.Where("gateway = ?", node.Gateway).First(&existing).Error; err == nil {
		//  the existing record
		existing.Status = node.Status
		return db.Save(&existing).Error
	}
	// Insert a new record
	return db.Create(node).Error
}

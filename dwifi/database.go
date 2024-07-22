package dwifi

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type NodeDwifi struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	Gateway        string    `json:"gateway"`
	Status         string    `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Password       string    `json:"password"`
	Location       string    `json:"location"`
	Price_per_min  string    `json:"price_per_min"`
	Wallet_Address string    `json:"wallet_address"`
}

func InitDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
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
		existing.Status = node.Status
		existing.Password = node.Password
		return db.Save(&existing).Error
	}
	return db.Create(node).Error
}

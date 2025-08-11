package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// New connects to the database and returns a *gorm.DB
func New(dsn string) (*gorm.DB, error) {
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("❌ Failed to connect to database:", err)
		return nil, err
	}

	if err := gormDB.AutoMigrate(&User{}); err != nil {
		log.Printf("auto migrate error: %v", err)
	}
	return gormDB, nil
}

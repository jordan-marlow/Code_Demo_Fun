package db

import (
	"POS/models"
	"log"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

// GetDB initializes or returns the existing DB connection
func GetDB() *gorm.DB {
	once.Do(func() {
		var err error
		db, err = gorm.Open(sqlite.Open("pos.db"), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		// AutoMigrate to keep the schema up to date
		db.AutoMigrate(&models.Product{}, &models.Order{})
	})
	return db
}

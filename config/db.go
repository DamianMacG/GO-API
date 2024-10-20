// config/db.go
package config

import (
	"GO-API/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=yourpassword dbname=albumsdb port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database!", err)
	}

	DB = database

	// Migrate the schema
	DB.AutoMigrate(&models.Album{}, &models.Review{}) // Add Review to the migration
}


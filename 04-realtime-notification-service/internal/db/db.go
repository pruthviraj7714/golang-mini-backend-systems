package db

import (
	"fmt"
	"realtime-notification-service/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(connString string) *gorm.DB {
	database, err := gorm.Open(postgres.Open(connString), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error while connecting to database: %w", err)
	}

	database.AutoMigrate(&models.Notification{})

	return database
}

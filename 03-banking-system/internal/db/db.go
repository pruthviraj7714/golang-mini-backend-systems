package db

import (
	"banking-system/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(connStr string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	db.AutoMigrate(&models.User{})

	return db

}

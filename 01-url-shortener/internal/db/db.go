package db

import (
	"fmt"
	"log"
	"url-shortener/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {

	db, err := gorm.Open(postgres.Open(config.LoadConfig().DBURL), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Database connected")

	return db
}

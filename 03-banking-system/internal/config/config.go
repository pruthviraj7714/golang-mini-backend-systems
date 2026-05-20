package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	JWTSECRET string
	DBURL     string
}

func LoadConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	return &Config{
		Port:      os.Getenv("PORT"),
		JWTSECRET: os.Getenv("JWT_SECRET"),
		DBURL:     os.Getenv("DB_URL"),
	}

}

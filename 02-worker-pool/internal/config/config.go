package config

import (
	"os"
	"strconv"
)

type Config struct {
	PORT             string
	DB_URL           string
	WORKER_POOL_SIZE int
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)

	if val == "" {
		return fallback
	}

	return val
}

func LoadConfig() *Config {

	workerSize, _ := strconv.Atoi(getEnv("WORKER_POOL_SIZE", "5"))

	return &Config{
		PORT:             getEnv("PORT", "8080"),
		DB_URL:           getEnv("DB_URL", "postgres://postgres:password@localhost:5432/postgres?sslmode=disable"),
		WORKER_POOL_SIZE: workerSize,
	}
}

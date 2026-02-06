package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	Env        string
	DBConn     string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, falling back to environment variables")
	}

	return &Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),
		Env:        getEnv("ENV", "development"),
		DBConn:     os.Getenv("DB_URL"), // Simplified for now
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort  string
	Env         string
	DatabaseURL string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, falling back to environment variables")
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		getEnv("BILLING_DB_USER", "user"),
		getEnv("BILLING_DB_PASSWORD", "password"),
		getEnv("BILLING_DB_HOST", "127.0.0.1"),
		getEnv("BILLING_DB_PORT", "5433"),
		getEnv("BILLING_DB_NAME", "billing_db"),
		getEnv("BILLING_DB_SSLMODE", "disable"),
	)

	return &Config{
		ServerPort:  getEnv("SERVER_PORT", "8080"),
		Env:         getEnv("ENV", "development"),
		DatabaseURL: dsn,
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

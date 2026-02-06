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
	RedisURL    string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, falling back to environment variables")
	}

	// Databases: Prioritize direct URLs if they exist (standard for Render/Cloud)
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			getEnv("BILLING_DB_USER", "user"),
			getEnv("BILLING_DB_PASSWORD", "password"),
			getEnv("BILLING_DB_HOST", "127.0.0.1"),
			getEnv("BILLING_DB_PORT", "5433"),
			getEnv("BILLING_DB_NAME", "billing_db"),
			getEnv("BILLING_DB_SSLMODE", "disable"),
		)
	}

	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		// Local fallback
		redisURL = fmt.Sprintf("redis://%s:%s",
			getEnv("REDIS_HOST", "localhost"),
			getEnv("REDIS_PORT", "6379"),
		)
	}

	return &Config{
		ServerPort:  getEnv("SERVER_PORT", "8080"),
		Env:         getEnv("ENV", "development"),
		DatabaseURL: dbURL,
		RedisURL:    redisURL,
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

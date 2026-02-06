package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/develoFavour/billing-engine-go/internal/config"
	"github.com/jackc/pgx/v5"
)

func main() {
	cfg := config.LoadConfig()
	fmt.Printf("Connecting to: %s\n", cfg.DatabaseURL)

	conn, err := pgx.Connect(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	migrationPath := filepath.Join("migrations", "000001_init_schema.up.sql")
	content, err := os.ReadFile(migrationPath)
	if err != nil {
		log.Fatalf("Unable to read migration file: %v\n", err)
	}

	fmt.Println("Running migration...")
	_, err = conn.Exec(context.Background(), string(content))
	if err != nil {
		log.Fatalf("Migration failed: %v\n", err)
	}

	fmt.Println("Migration successful!")
}

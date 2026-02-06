package main

import (
	"context"
	"fmt"
	"log"

	"github.com/develoFavour/billing-engine-go/internal/config"
	"github.com/develoFavour/billing-engine-go/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func main() {
	cfg := config.LoadConfig()

	conn, err := pgx.Connect(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	fmt.Println("Seeding initial pricing tiers and customer...")

	// 1. Create a Pricing Tier
	tierID := uuid.New()
	_, err = conn.Exec(context.Background(), `
		INSERT INTO pricing_tiers (id, name, resource_type, model, unit_price)
		VALUES ($1, $2, $3, $4, $5)
	`, tierID, "Pro Plan", models.ResourceTypeAPICall, "flat", 0.01)
	if err != nil {
		log.Fatalf("Failed to seed pricing tier: %v\n", err)
	}

	// 2. Create a Customer
	customerID := uuid.New()
	_, err = conn.Exec(context.Background(), `
		INSERT INTO customers (id, email, name, pricing_tier_id)
		VALUES ($1, $2, $3, $4)
	`, customerID, "test@example.com", "Test Corp", tierID)
	if err != nil {
		log.Fatalf("Failed to seed customer: %v\n", err)
	}

	fmt.Printf("\nSeed Successful!\n")
	fmt.Printf("Customer ID: %s\n", customerID)
	fmt.Printf("Pricing Tier: %s ($0.01 per API call)\n", tierID)
}

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/develoFavour/billing-engine-go/internal/config"
	"github.com/jackc/pgx/v5"
)

func main() {
	cfg := config.LoadConfig()
	conn, err := pgx.Connect(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	var id string
	err = conn.QueryRow(context.Background(), "SELECT id FROM customers LIMIT 1").Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("CUSTOMER_ID=%s\n", id)
}

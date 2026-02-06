package service

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/develoFavour/billing-engine-go/internal/models"
	"github.com/develoFavour/billing-engine-go/internal/repository"
	"github.com/google/uuid"
)

type BillingService interface {
	GetEstimatedBill(ctx context.Context, customerID string) (float64, error)
	AggregateUsage(ctx context.Context) error
}

type billingService struct {
	usageRepo repository.UsageRepository
	meterRepo repository.MeterRepository
	// We'll need a way to get pricing, for now let's assume we pass a DB pool or specific repo
	// For speed, let's add a simplify pricing check
}

func NewBillingService(usageRepo repository.UsageRepository, meterRepo repository.MeterRepository) BillingService {
	return &billingService{
		usageRepo: usageRepo,
		meterRepo: meterRepo,
	}
}

func (s *billingService) GetEstimatedBill(ctx context.Context, customerID string) (float64, error) {
	// 1. Get real-time usage from Redis
	// For this demo, let's check API calls
	usage, err := s.meterRepo.GetTotalUsage(ctx, customerID, models.ResourceTypeAPICall)
	if err != nil {
		return 0, err
	}

	// 2. Apply pricing logic (Hardcoded for POC, would normally use a repository)
	pricePerUnit := 0.01 // $0.01 per call
	return usage * pricePerUnit, nil
}

func (s *billingService) AggregateUsage(ctx context.Context) error {
	log.Println("Starting background usage aggregation...")

	// 1. Find all usage keys in Redis
	keys, err := s.meterRepo.ScanKeys(ctx, "usage:*")
	if err != nil {
		return fmt.Errorf("failed to scan usage keys: %w", err)
	}

	for _, key := range keys {
		// Key format: usage:customer_id:resource_type
		parts := strings.Split(key, ":")
		if len(parts) != 3 {
			continue
		}

		customerIDStr := parts[1]
		resourceType := models.ResourceType(parts[2])

		customerID, err := uuid.Parse(customerIDStr)
		if err != nil {
			log.Printf("Invalid customer ID in Redis key %s: %v", key, err)
			continue
		}

		// 2. Get the value
		quantity, err := s.meterRepo.GetTotalUsage(ctx, customerIDStr, resourceType)
		if err != nil {
			log.Printf("Failed to get usage for key %s: %v", key, err)
			continue
		}

		if quantity == 0 {
			continue
		}

		// 3. Persist to Postgres
		event := &models.UsageEvent{
			ID:           uuid.New(),
			CustomerID:   customerID,
			ResourceType: resourceType,
			Quantity:     quantity,
			Timestamp:    time.Now().UTC(),
		}

		if err := s.usageRepo.Create(ctx, event); err != nil {
			log.Printf("Failed to persist usage for customer %s: %v", customerID, err)
			continue
		}

		// 4. Reset Redis (Clear the meter)
		if err := s.meterRepo.ResetUsage(ctx, customerIDStr, resourceType); err != nil {
			log.Printf("Failed to reset usage for customer %s in Redis: %v", customerID, err)
		}

		log.Printf("Aggregated %.2f units for customer %s", quantity, customerID)
	}

	return nil
}

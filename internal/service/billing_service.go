package service

import (
	"context"
	"log"

	"github.com/develoFavour/billing-engine-go/internal/models"
	"github.com/develoFavour/billing-engine-go/internal/repository"
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
	// Logic to scan all active Redis keys and "flush" them to Postgres
	// This is a more advanced step we'll build in the backgrounds worker
	return nil
}

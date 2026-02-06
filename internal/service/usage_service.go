package service

import (
	"context"
	"time"

	"github.com/develoFavour/billing-engine-go/internal/models"
	"github.com/develoFavour/billing-engine-go/internal/repository"
	"github.com/google/uuid"
)

type UsageService interface {
	RecordUsage(ctx context.Context, customerID uuid.UUID, resourceType models.ResourceType, quantity float64, metadata models.Metadata) (*models.UsageEvent, error)
	GetCustomerUsage(ctx context.Context, customerID string) ([]models.UsageEvent, error)
}

type usageService struct {
	usageRepo repository.UsageRepository
	meterRepo repository.MeterRepository
}

func NewUsageService(usageRepo repository.UsageRepository, meterRepo repository.MeterRepository) UsageService {
	return &usageService{
		usageRepo: usageRepo,
		meterRepo: meterRepo,
	}
}

func (s *usageService) RecordUsage(ctx context.Context, customerID uuid.UUID, resourceType models.ResourceType, quantity float64, metadata models.Metadata) (*models.UsageEvent, error) {
	// High-performance write to Redis
	if err := s.meterRepo.IncrementUsage(ctx, customerID.String(), resourceType, quantity); err != nil {
		return nil, err
	}

	// For auditing, we still return a model (but we don't block the API waiting for Postgres)
	return &models.UsageEvent{
		ID:           uuid.New(),
		CustomerID:   customerID,
		ResourceType: resourceType,
		Quantity:     quantity,
		Timestamp:    time.Now().UTC(),
		Metadata:     metadata,
	}, nil
}

func (s *usageService) GetCustomerUsage(ctx context.Context, customerID string) ([]models.UsageEvent, error) {
	return s.usageRepo.GetByCustomerID(ctx, customerID)
}

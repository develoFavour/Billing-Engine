package repository

import (
	"context"

	"github.com/develoFavour/billing-engine-go/internal/models"
)

type MeterRepository interface {
	IncrementUsage(ctx context.Context, customerID string, resourceType models.ResourceType, quantity float64) error
	GetTotalUsage(ctx context.Context, customerID string, resourceType models.ResourceType) (float64, error)
	ResetUsage(ctx context.Context, customerID string, resourceType models.ResourceType) error
	ScanKeys(ctx context.Context, pattern string) ([]string, error)
}

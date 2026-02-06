package repository

import (
	"context"

	"github.com/develoFavour/billing-engine-go/internal/models"
)

type UsageRepository interface {
	Create(ctx context.Context, event *models.UsageEvent) error
	GetByCustomerID(ctx context.Context, customerID string) ([]models.UsageEvent, error)
}

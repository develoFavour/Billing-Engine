package models

import (
	"time"

	"github.com/google/uuid"
)

type PricingModel string

const (
	PricingModelFlat   PricingModel = "flat"
	PricingModelTiered PricingModel = "tiered"
)

type Customer struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	TierID    uuid.UUID `json:"tier_id"`
	CreatedAt time.Time `json:"created_at"`
}

type PricingTier struct {
	ID           uuid.UUID    `json:"id"`
	Name         string       `json:"name"`
	ResourceType ResourceType `json:"resource_type"`
	Model        PricingModel `json:"model"`
	UnitPrice    float64      `json:"unit_price"` // For flat pricing
	UpdatedAt    time.Time    `json:"updated_at"`
}

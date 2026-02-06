package models

import (
	"time"

	"github.com/google/uuid"
)

type ResourceType string

const (
	ResourceTypeAPICall   ResourceType = "api_call"
	ResourceTypeStorageGB ResourceType = "storage_gb"
	ResourceTypeCompute   ResourceType = "compute_instance"
)

// UsageEvent represents a single metric event reported by a system.
type UsageEvent struct {
	ID           uuid.UUID    `json:"id"`
	CustomerID   uuid.UUID    `json:"customer_id"`
	ResourceType ResourceType `json:"resource_type"`
	Quantity     float64      `json:"quantity"`
	Timestamp    time.Time    `json:"timestamp"`
	Metadata     Metadata     `json:"metadata,omitempty"`
}

type Metadata map[string]interface{}

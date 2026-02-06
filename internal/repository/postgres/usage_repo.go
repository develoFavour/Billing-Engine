package postgres

import (
	"context"
	"fmt"

	"github.com/develoFavour/billing-engine-go/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type usageRepository struct {
	db *pgxpool.Pool
}

func NewUsageRepository(db *pgxpool.Pool) *usageRepository {
	return &usageRepository{db: db}
}

func (r *usageRepository) Create(ctx context.Context, event *models.UsageEvent) error {
	query := `
		INSERT INTO usage_records (id, customer_id, resource_type, quantity, timestamp, metadata)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.db.Exec(ctx, query,
		event.ID,
		event.CustomerID,
		event.ResourceType,
		event.Quantity,
		event.Timestamp,
		event.Metadata,
	)
	if err != nil {
		return fmt.Errorf("failed to insert usage record: %w", err)
	}
	return nil
}

func (r *usageRepository) GetByCustomerID(ctx context.Context, customerID string) ([]models.UsageEvent, error) {
	query := `
		SELECT id, customer_id, resource_type, quantity, timestamp, metadata
		FROM usage_records
		WHERE customer_id = $1
		ORDER BY timestamp DESC
	`
	rows, err := r.db.Query(ctx, query, customerID)
	if err != nil {
		return nil, fmt.Errorf("failed to query usage records: %w", err)
	}
	defer rows.Close()

	var events []models.UsageEvent
	for rows.Next() {
		var event models.UsageEvent
		err := rows.Scan(
			&event.ID,
			&event.CustomerID,
			&event.ResourceType,
			&event.Quantity,
			&event.Timestamp,
			&event.Metadata,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan usage record: %w", err)
		}
		events = append(events, event)
	}

	return events, nil
}

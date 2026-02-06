package redis

import (
	"context"
	"fmt"
	"strconv"

	"github.com/develoFavour/billing-engine-go/internal/models"
	"github.com/redis/go-redis/v9"
)

type meterRepository struct {
	rdb *redis.Client
}

func NewMeterRepository(rdb *redis.Client) *meterRepository {
	return &meterRepository{rdb: rdb}
}

func (r *meterRepository) IncrementUsage(ctx context.Context, customerID string, resourceType models.ResourceType, quantity float64) error {
	key := fmt.Sprintf("usage:%s:%s", customerID, resourceType)

	// Atomic increment
	err := r.rdb.IncrByFloat(ctx, key, quantity).Err()
	if err != nil {
		return fmt.Errorf("failed to increment redis usage: %w", err)
	}
	return nil
}

func (r *meterRepository) GetTotalUsage(ctx context.Context, customerID string, resourceType models.ResourceType) (float64, error) {
	key := fmt.Sprintf("usage:%s:%s", customerID, resourceType)

	val, err := r.rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(val, 64)
}

func (r *meterRepository) ResetUsage(ctx context.Context, customerID string, resourceType models.ResourceType) error {
	key := fmt.Sprintf("usage:%s:%s", customerID, resourceType)
	return r.rdb.Del(ctx, key).Err()
}

package worker

import (
	"context"
	"log"
	"time"

	"github.com/develoFavour/billing-engine-go/internal/service"
)

type AggregatorWorker struct {
	billingService service.BillingService
	interval       time.Duration
}

func NewAggregatorWorker(billingService service.BillingService, interval time.Duration) *AggregatorWorker {
	return &AggregatorWorker{
		billingService: billingService,
		interval:       interval,
	}
}

func (w *AggregatorWorker) Start(ctx context.Context) {
	log.Printf("Starting Aggregator Worker with interval %v", w.interval)

	ticker := time.NewTicker(w.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := w.billingService.AggregateUsage(ctx); err != nil {
				log.Printf("Aggregator worker error: %v", err)
			}
		case <-ctx.Done():
			log.Println("Stopping Aggregator Worker...")
			return
		}
	}
}

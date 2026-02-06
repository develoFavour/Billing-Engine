package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/develoFavour/billing-engine-go/internal/api/handlers"
	"github.com/develoFavour/billing-engine-go/internal/api/routes"
	"github.com/develoFavour/billing-engine-go/internal/config"
	"github.com/develoFavour/billing-engine-go/internal/repository/postgres"
	"github.com/develoFavour/billing-engine-go/internal/repository/redis"
	"github.com/develoFavour/billing-engine-go/internal/service"
	"github.com/develoFavour/billing-engine-go/internal/worker"
	"github.com/develoFavour/billing-engine-go/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// 1. Load Configuration
	cfg := config.LoadConfig()

	// 2. Initialize Database Connections
	// postgres
	pool, err := database.NewPostgresPool(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to initialize postgres: %v", err)
	}
	defer pool.Close()

	// redis
	rdb, err := database.NewRedisClient("127.0.0.1", "6379", "", 0)
	if err != nil {
		log.Fatalf("failed to initialize redis: %v", err)
	}
	defer rdb.Close()

	// 3. Initialize Repositories (Dependency Injection)
	usageRepo := postgres.NewUsageRepository(pool)
	meterRepo := redis.NewMeterRepository(rdb)

	// 4. Initialize Services (Adding MeterRepo to UsageService)
	usageService := service.NewUsageService(usageRepo, meterRepo)
	usageHandler := handlers.NewUsageHandler(usageService)

	billingService := service.NewBillingService(usageRepo, meterRepo)
	billingHandler := handlers.NewBillingHandler(billingService)

	// 5. Initialize & Start Background Workers
	aggregator := worker.NewAggregatorWorker(billingService, 1*time.Minute)
	go aggregator.Start(ctx)

	router := gin.Default()

	// Basic Health Check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
			"time":   time.Now().Format(time.RFC3339),
			"config": cfg.Env,
		})
	})

	// Setup API Routes
	routes.SetupRoutes(router, usageHandler, billingHandler)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.ServerPort),
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Println("Server started on :8080")

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}

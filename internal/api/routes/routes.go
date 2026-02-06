package routes

import (
	"github.com/develoFavour/billing-engine-go/internal/api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, usageHandler *handlers.UsageHandler) {
	v1 := router.Group("/api/v1")
	{
		usage := v1.Group("/usage")
		{
			usage.POST("", usageHandler.RecordUsage)
			usage.GET("/:customer_id", usageHandler.GetUsage)
		}
	}
}

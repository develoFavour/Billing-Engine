package handlers

import (
	"net/http"

	"github.com/develoFavour/billing-engine-go/internal/models"
	"github.com/develoFavour/billing-engine-go/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UsageHandler struct {
	service service.UsageService
}

func NewUsageHandler(service service.UsageService) *UsageHandler {
	return &UsageHandler{service: service}
}

type RecordUsageRequest struct {
	CustomerID   string              `json:"customer_id" binding:"required,uuid"`
	ResourceType models.ResourceType `json:"resource_type" binding:"required"`
	Quantity     float64             `json:"quantity" binding:"required,gt=0"`
	Metadata     models.Metadata     `json:"metadata"`
}

func (h *UsageHandler) RecordUsage(c *gin.Context) {
	var req RecordUsageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customerUUID, _ := uuid.Parse(req.CustomerID)

	event, err := h.service.RecordUsage(c.Request.Context(), customerUUID, req.ResourceType, req.Quantity, req.Metadata)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to record usage"})
		return
	}

	c.JSON(http.StatusCreated, event)
}

func (h *UsageHandler) GetUsage(c *gin.Context) {
	customerID := c.Param("customer_id")
	if _, err := uuid.Parse(customerID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid customer_id"})
		return
	}

	events, err := h.service.GetCustomerUsage(c.Request.Context(), customerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch usage records"})
		return
	}

	c.JSON(http.StatusOK, events)
}

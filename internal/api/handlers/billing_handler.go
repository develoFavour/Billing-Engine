package handlers

import (
	"net/http"

	"github.com/develoFavour/billing-engine-go/internal/service"
	"github.com/gin-gonic/gin"
)

type BillingHandler struct {
	service service.BillingService
}

func NewBillingHandler(service service.BillingService) *BillingHandler {
	return &BillingHandler{service: service}
}

func (h *BillingHandler) GetEstimatedBill(c *gin.Context) {
	customerID := c.Param("customer_id")

	bill, err := h.service.GetEstimatedBill(c.Request.Context(), customerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to calculate bill"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"customer_id":    customerID,
		"currency":       "USD",
		"estimated_bill": bill,
	})
}

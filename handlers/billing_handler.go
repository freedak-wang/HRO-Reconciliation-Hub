package handlers

import (
	"github.com/gin-gonic/gin"
	"your-project/services"
	"net/http"
	"time"
)

type BillingHandler struct {
	billingService *services.BillingService
}

func NewBillingHandler(billingService *services.BillingService) *BillingHandler {
	return &BillingHandler{billingService: billingService}
}

func (h *BillingHandler) GenerateMonthlyBill(c *gin.Context) {
	var req struct {
		ContractID uint      `json:"contract_id" binding:"required"`
		BillMonth  time.Time `json:"bill_month" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bill, err := h.billingService.GenerateMonthlyBill(req.ContractID, req.BillMonth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bill)
}

func (h *BillingHandler) CalculateDifferences(c *gin.Context) {
	billID := c.Param("billID")

	differences, err := h.billingService.CalculateDifferences(billID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, differences)
}

func (h *BillingHandler) ApplyAdjustments(c *gin.Context) {
	var req struct {
		BillID      uint                       `json:"bill_id" binding:"required"`
		Adjustments []models.DifferenceRecord `json:"adjustments" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.billingService.ApplyAdjustments(req.BillID, req.Adjustments)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Adjustments applied successfully"})
}
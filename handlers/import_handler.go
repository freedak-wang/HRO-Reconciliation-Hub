package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"your-project/services"
)

type ImportHandler struct {
	importService *services.ImportService
}

func NewImportHandler(importService *services.ImportService) *ImportHandler {
	return &ImportHandler{importService: importService}
}

func (h *ImportHandler) ImportCustomerBills(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	err = h.importService.ImportCustomerBills(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer bills imported successfully"})
}

func (h *ImportHandler) ImportActualPayments(c *gin.Context) {
	// Similar implementation to ImportCustomerBills
}

func (h *ImportHandler) ImportSalesInvoices(c *gin.Context) {
	// Similar implementation to ImportCustomerBills
}

func (h *ImportHandler) ImportSocialInsurancePayments(c *gin.Context) {
	// Similar implementation to ImportCustomerBills
}

func (h *ImportHandler) ImportSalaryTaxPayments(c *gin.Context) {
	// Similar implementation to ImportCustomerBills
}

func (h *ImportHandler) ImportOtherPayments(c *gin.Context) {
	// Similar implementation to ImportCustomerBills
}
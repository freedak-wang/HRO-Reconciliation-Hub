package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"your-project/handlers"
	"your-project/services"
)

func main() {
	// TODO: Initialize database connection
	var db *gorm.DB

	billingService := services.NewBillingService(db)
	contractService := services.NewContractService(db)
	personnelService := services.NewPersonnelService(db)
	importService := services.NewImportService(db)

	billingHandler := handlers.NewBillingHandler(billingService)
	contractHandler := handlers.NewContractHandler(contractService)
	personnelHandler := handlers.NewPersonnelHandler(personnelService)
	importHandler := handlers.NewImportHandler(importService)

	r := gin.Default()

	// TODO: Add authentication middleware

	api := r.Group("/api")
	{
		// Existing routes...

		// Import routes
		api.POST("/import/customer-bills", importHandler.ImportCustomerBills)
		api.POST("/import/actual-payments", importHandler.ImportActualPayments)
		api.POST("/import/sales-invoices", importHandler.ImportSalesInvoices)
		api.POST("/import/social-insurance-payments", importHandler.ImportSocialInsurancePayments)
		api.POST("/import/salary-tax-payments", importHandler.ImportSalaryTaxPayments)
		api.POST("/import/other-payments", importHandler.ImportOtherPayments)
	}

	r.Run(":8080")
}
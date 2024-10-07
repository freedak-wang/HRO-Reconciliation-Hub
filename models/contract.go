package models

import (
	"gorm.io/gorm"
	"time"
)

type Contract struct {
	gorm.Model
	ClientID       uint
	ClientName     string
	StartDate      time.Time
	EndDate        time.Time
	ExpectedAmount float64
	Status         string // "active", "expired", "terminated"
	ServiceItems   []ServiceItem
	MonthlyBills   []MonthlyBill
}

type ServiceItem struct {
	gorm.Model
	ContractID uint
	Name       string
	UnitPrice  float64
	Quantity   int
}
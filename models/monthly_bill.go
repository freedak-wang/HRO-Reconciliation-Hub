package models

import (
	"gorm.io/gorm"
	"time"
)

type MonthlyBill struct {
	gorm.Model
	ContractID        uint
	BillMonth         time.Time
	TotalAmount       float64
	AdjustmentAmount  float64
	AdjustmentReason  string
	PersonnelServices []PersonnelService
	NonPersonalItems  []NonPersonalItem
}

type PersonnelService struct {
	gorm.Model
	MonthlyBillID uint
	PersonnelID   uint
	ServiceItemID uint
	Amount        float64
}

type NonPersonalItem struct {
	gorm.Model
	MonthlyBillID uint
	Description   string
	Amount        float64
}
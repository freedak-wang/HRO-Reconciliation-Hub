package models

import (
	"gorm.io/gorm"
	"time"
)

type DifferenceRecord struct {
	gorm.Model
	ContractID    uint
	MonthlyBillID uint
	DiffType      string // "personnel_change", "service_change", "amount_diff", "non_personal"
	Description   string
	Amount        float64
	Status        string // "pending", "adjusted", "confirmed"
	ProcessNote   string
}

type PersonnelChangeRecord struct {
	gorm.Model
	ContractID  uint
	ChangeMonth time.Time
	ChangeType  string // "add", "remove"
	PersonnelID uint
	ChangeDate  time.Time
	ServiceItems []uint // IDs of related service items
}
package models

import (
	"gorm.io/gorm"
	"time"
)

type Personnel struct {
	gorm.Model
	Name       string
	ContractID uint
	StartDate  time.Time
	EndDate    *time.Time
	Status     string // "active", "inactive"
}

type PersonnelChange struct {
	gorm.Model
	PersonnelID uint
	ContractID  uint
	ChangeType  string // "add", "remove"
	ChangeDate  time.Time
	Reason      string
}
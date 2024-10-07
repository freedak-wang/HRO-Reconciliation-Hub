package models

import (
	"gorm.io/gorm"
	"time"
)

type CustomerBill struct {
	gorm.Model
	ContractID     uint
	BillMonth      time.Time
	TotalAmount    float64
	BillItems      []BillItem
	Status         string // "draft", "confirmed", "paid"
}

type BillItem struct {
	gorm.Model
	CustomerBillID uint
	PersonnelID    uint
	ServiceItemID  uint
	Quantity       int
	UnitPrice      float64
	Amount         float64
}

type ActualPayment struct {
	gorm.Model
	ContractID     uint
	PaymentDate    time.Time
	Amount         float64
	PaymentMethod  string
	Reference      string
}

type SalesInvoice struct {
	gorm.Model
	ContractID     uint
	InvoiceDate    time.Time
	InvoiceNumber  string
	Amount         float64
	Status         string // "issued", "paid", "cancelled"
}

type SocialInsurancePayment struct {
	gorm.Model
	ContractID     uint
	PaymentMonth   time.Time
	TotalAmount    float64
	EmployerPart   float64
	EmployeePart   float64
	PersonnelCount int
}

type SalaryTaxPayment struct {
	gorm.Model
	ContractID     uint
	PaymentMonth   time.Time
	TotalSalary    float64
	TaxAmount      float64
	PersonnelCount int
}

type OtherPayment struct {
	gorm.Model
	ContractID     uint
	PaymentDate    time.Time
	Amount         float64
	Description    string
	Category       string
}
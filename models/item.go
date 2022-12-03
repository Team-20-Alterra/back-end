package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name        string `validate:"required" json:"name" form:"name"`
	Amount      uint64 `validate:"required" json:"amount" form:"amount"`
	UnitPrice   uint64 	`validate:"required" json:"unit_price" form:"unit_price"`
	TotalPrice  uint64 `validate:"required" json:"total_price" form:"total_price"`
	InvoiceID 	int `validate:"required" json:"invoice_id" form:"invoice_id"`
	Invoice     Invoice
}


type ItemResponse struct {
	Name        string `validate:"required" json:"name" form:"name"`
	Amount      uint64 `validate:"required" json:"amount" form:"amount"`
	UnitPrice   uint64 	`validate:"required" json:"unit_price" form:"unit_price"`
	TotalPrice  uint64 `validate:"required" json:"total_price" form:"total_price"`
	InvoiceID 	int `validate:"required" json:"invoice_id" form:"invoice_id"`
}
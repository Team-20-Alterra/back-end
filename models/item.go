package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name        string `validate:"required" json:"name" form:"name"`
	Amount      string `gorm:"type:text" validate:"required" json:"amount" form:"amount"`
	UnitPrice   bool 	`validate:"required" json:"unit_price" form:"unit_price"`
	TotalPrice  string `validate:"required" json:"total_price" form:"total_price"`
	InvoiceID 	uint `validate:"required" json:"customer_id" form:"customer_id"`
	Invoice     Invoice
}

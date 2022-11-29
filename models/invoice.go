package models

import (
	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	Date    string `validate:"required" json:"date" form:"date"`
	NoInvoice string `validate:"required" json:"no_invoice" form:"no_invoice"`
	Price   string `validate:"required" json:"price" form:"price"`
	Payment string `validate:"required" json:"payment" form:"payment"`
	Type    string `validate:"required" json:"type" form:"type"`
	Status  string `validate:"required" json:"status" form:"status"`
	Total   string `validate:"required" json:"total" form:"total"`
	Subtotal string `validate:"required" json:"sub_total" form:"sub_total"`
	BusinnesID int	`validate:"required" json:"businnes_id" form:"businnes_id"`
	UserID  int    `json:"user_id" form:"user_id"`
	User    User
}
type InvoiceResponse struct {
	Date    string `validate:"required" json:"date" form:"date"`
	Price   string `validate:"required" json:"price" form:"price"`
	Payment string `validate:"required" json:"payment" form:"payment"`
	Type    string `validate:"required" json:"type" form:"type"`
	Status  string `validate:"required" json:"status" form:"status"`
	UserID  int    `json:"user_id" form:"user_id"`
}

type InvoiceStatus struct {
	Status  string `validate:"required" json:"status" form:"status"`
}
package models

import (
	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	DatePay       string `validate:"required" json:"date_pay" form:"date_pay"`
	BillingDate   string `validate:"required" json:"billing_date" form:"billing_date"`
	NoInvoice     string `validate:"required" json:"no_invoice" form:"no_invoice"`
	Price         string `validate:"required" json:"price" form:"price"`
	Payment       string `validate:"required" json:"payment" form:"payment"`
	Type          string `validate:"required" json:"type" form:"type"`
	StatusInvoice string `json:"status_invoice" form:"status_invoice"`
	Status        string `validate:"required" json:"status" form:"status"`
	Total         string `validate:"required" json:"total" form:"total"`
	Subtotal      string `validate:"required" json:"sub_total" form:"sub_total"`
	BusinnesID    int    `validate:"required" json:"businnes_id" form:"businnes_id"`
	UserID        int    `json:"user_id" form:"user_id"`
	User          User
	// Business      Business
	Item          []Item
}
type InvoiceResponse struct {
	DatePay       string `validate:"required" json:"date_pay" form:"date_pay"`
	BillingDate   string `validate:"required" json:"billing_date" form:"billing_date"`
	Price   string `validate:"required" json:"price" form:"price"`
	Payment string `validate:"required" json:"payment" form:"payment"`
	Type    string `validate:"required" json:"type" form:"type"`
	Status  string `validate:"required" json:"status" form:"status"`
	UserID  int    `json:"user_id" form:"user_id"`
	BusinnesID    int    `validate:"required" json:"businnes_id" form:"businnes_id"`
}

type InvoiceUpdate struct {
	NoInvoice     string `validate:"required" json:"no_invoice" form:"no_invoice"`
	Price   string `validate:"required" json:"price" form:"price"`
	Payment string `validate:"required" json:"payment" form:"payment"`
	Type    string `validate:"required" json:"type" form:"type"`
	Status  string `validate:"required" json:"status" form:"status"`
	UserID  int    `json:"user_id" form:"user_id"`
}

type InvoicePembayaranStatus struct {
	DatePay       string `validate:"required" json:"date_pay" form:"date_pay"`
	Payment string `validate:"required" json:"payment" form:"payment"`
	Status  string `validate:"required" json:"status" form:"status"`
}
type InvoiceStatus struct {
	StatusInvoice string `validate:"required" json:"status_invoice" form:"status_invoice"`
}

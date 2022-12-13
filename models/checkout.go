package models

import "gorm.io/gorm"

type Checkout struct {
	gorm.Model
	DatePay       string `validate:"required" json:"date_pay" form:"date_pay"`
	BillingDate   string `validate:"required" json:"billing_date" form:"billing_date"`
	ListBankID    int    `json:"list_id" form:"list_id"`
	InvoiceID 	  int    `validate:"required" json:"invoice_id" form:"invoice_id"`
	ListBank      ListBank
}
type CheckoutInput struct {
	BillingDate   string `validate:"required" json:"billing_date" form:"billing_date"`
	ListBankID    int    `json:"list_id" form:"list_id"`
	InvoiceID 	  int    `validate:"required" json:"invoice_id" form:"invoice_id"`
}
type CheckoutUpdate struct {
	DatePay       string `validate:"required" json:"date_pay" form:"date_pay"`
}



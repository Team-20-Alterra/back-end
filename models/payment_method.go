package models

import "gorm.io/gorm"

type PaymentMethod struct {
	gorm.Model
	Body   string `json:"body" form:"body"`
	BankID int    `json:"bank_id" form:"bank_id"`
	Bank   Bank
}

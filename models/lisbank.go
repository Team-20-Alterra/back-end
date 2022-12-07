package models

import "gorm.io/gorm"

type ListBank struct {
	gorm.Model
	Owner string `validate:"required" json:"owner" form:"owner"`
	AccountNumber string `validate:"required" json:"account_number" form:"account_number"`
	BankID  int    `json:"bank_id" form:"bank_id"`
	BusinnesID 	int `validate:"required" json:"businnes_id" form:"businnes_id"`
	Bank    Bank
	Businnes    Business
}


type LisBankInput struct {
	Owner string `validate:"required" json:"owner" form:"owner"`
	AccountNumber string `validate:"required" json:"account_number" form:"account_number"`
	BankID  int    `json:"bank_id" form:"bank_id"`
	BusinnesID 	int `validate:"required" json:"businnes_id" form:"businnes_id"`
}
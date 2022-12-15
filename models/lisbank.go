package models

import "gorm.io/gorm"

type ListBank struct {
	gorm.Model
	Owner         string `validate:"required" json:"owner" form:"owner"`
	AccountNumber string `validate:"required" json:"account_number" form:"account_number"`
	BankID        int    `json:"bank_id" form:"bank_id"`
	BusinessID 	  int    `validate:"required" json:"business_id" form:"business_id"`
	// Business      Business
	Bank          Bank
}


type LisBankResponse struct {
	ID      int `validate:"required" json:"id" form:"id"`
	Owner string `validate:"required" json:"owner" form:"owner"`
	AccountNumber string `validate:"required" json:"account_number" form:"account_number"`
	BankID  int    `json:"bank_id" form:"bank_id"`
	BusinessID 	int `json:"business_id" form:"business_id"`
	Bank          BankResponseFK
	// Business      BusinessResponseFK
}
type LisBankInput struct {
	Owner string `validate:"required" json:"owner" form:"owner"`
	AccountNumber string `validate:"required" json:"account_number" form:"account_number"`
	BankID  int    `json:"bank_id" form:"bank_id"`
	BusinessID 	int `json:"business_id" form:"business_id"`
}
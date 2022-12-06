package models

import "gorm.io/gorm"

type Business struct {
	gorm.Model
	Name    string `validate:"required" json:"name" form:"name"`
	Address string `validate:"required" json:"address" form:"address"`
	No_telp string `validate:"required" json:"no_telp" form:"no_telp"`
	Type    string `validate:"required" json:"type" form:"type"`
	Logo    string `json:"logo" form:"logo"`
	UserID  int	   `json:"user_id" form:"user_id"`
	User    User
}

type BusinessInput struct {
	Name    string `validate:"required" json:"name" form:"name"`
	Address string `validate:"required" json:"address" form:"address"`
	No_telp string `validate:"required" json:"no_telp" form:"no_telp"`
	Type    string `validate:"required" json:"type" form:"type"`
	Logo    string `json:"logo" form:"logo"`
	UserID  int    `json:"user_id" form:"user_id"`
	// Owner string `validate:"required" json:"owner" form:"owner"`
	// AccountNumber string `validate:"required" json:"account_number" form:"account_number"`
	// BankID  int    `json:"bank_id" form:"bank_id"`
	// BusinnesID 	int `validate:"required" json:"businnes_id" form:"businnes_id"`
}

package models

import "gorm.io/gorm"

type Business struct {
	gorm.Model
	Name    string `validate:"required" json:"name" form:"name"`
	Address string `validate:"required" json:"address" form:"address"`
	No_telp string `validate:"required" json:"no_telp" form:"no_telp"`
	Type    string `validate:"required" json:"type" form:"type"`
	Logo    string `json:"logo" form:"logo"`
	BankID  int    `json:"bank_id" form:"bank_id"`
	Bank    Bank
}

type BusinessInput struct {
	Name    string `validate:"required" json:"name" form:"name"`
	Address string `validate:"required" json:"address" form:"address"`
	No_telp string `validate:"required" json:"no_telp" form:"no_telp"`
	Type    string `validate:"required" json:"type" form:"type"`
	Logo    string `json:"logo" form:"logo"`
	BankID  int    `json:"bank_id" form:"bank_id"`
}

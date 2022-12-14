package models

import "gorm.io/gorm"

type Bank struct {
	gorm.Model
	Name string `validate:"required" json:"name" form:"name"`
	Code string `validate:"required" json:"code" form:"code"`
	Logo string `json:"logo" form:"logo"`
}
type BankResponseFK struct {
	ID      int `validate:"required" json:"id" form:"id"`
	Name string `validate:"required" json:"name" form:"name"`
	Code string `validate:"required" json:"code" form:"code"`
	Logo string `json:"logo" form:"logo"`
}

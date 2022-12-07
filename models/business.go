package models

import "gorm.io/gorm"

type Business struct {
	gorm.Model
	Name    string `validate:"required" json:"name" form:"name"`
	Email   string `validate:"required" json:"email" form:"email"`
	Address string `validate:"required" json:"address" form:"address"`
	No_telp string `validate:"required" json:"no_telp" form:"no_telp"`
	Type    string `validate:"required" json:"type" form:"type"`
	Logo    string `json:"logo" form:"logo"`
	Reminder string `json:"reminder" form:"reminder"`
	Due_Date  string `json:"due_date" from:"due_date"`
	UserID  int `json:"user_id" form:"user_id"`
	User    User
}

type BusinessInput struct {
	Name    string `validate:"required" json:"name" form:"name"`
	Address string `validate:"required" json:"address" form:"address"`
	No_telp string `validate:"required" json:"no_telp" form:"no_telp"`
	Type    string `validate:"required" json:"type" form:"type"`
	Logo    string `json:"logo" form:"logo"`
	BankID  int    `validate:"required" json:"bank_id" form:"bank_id"`
	UserID  int    `json:"user_id" form:"user_id"`
}

type BusinessUpdate struct {
	Name    string `validate:"required" json:"name" form:"name"`
	Email   string `validate:"required" json:"email" form:"email"`
	Address string `validate:"required" json:"address" form:"address"`
	No_telp string `validate:"required" json:"no_telp" form:"no_telp"`
	Type    string `validate:"required" json:"type" form:"type"`
	Reminder string `validate:"required" json:"reminder" form:"reminder"`
	Due_Date  string `validate:"required" json:"due_date" from:"due_date"`
}

type BusinessLogo struct {
	Logo    string `json:"logo" form:"logo"`
}

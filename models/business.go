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
	// Reminder int64 `json:"reminder" form:"reminder"`
	// Due_Date  int64 `json:"due_date" form:"due_date"`
	UserID  int `json:"user_id" form:"user_id"`
	User    User
}
type BusinessResponse struct {
	Name    string `validate:"required" json:"name" form:"name"`
	Email   string `validate:"required" json:"email" form:"email"`
	Address string `validate:"required" json:"address" form:"address"`
	No_telp string `validate:"required" json:"no_telp" form:"no_telp"`
	Type    string `validate:"required" json:"type" form:"type"`
	Logo    string `json:"logo" form:"logo"`
	UserID  int `json:"user_id" form:"user_id"`
}

type BusinessInput struct {
	Name    string `validate:"required" json:"name" form:"name"`
	Email   string `validate:"required" json:"email" form:"email"`
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
	Reminder int64 `validate:"required" json:"reminder" form:"reminder"`
	Due_Date  int64 `validate:"required" json:"due_date" form:"due_date"`
	Logo    string `json:"logo" form:"logo"`
}


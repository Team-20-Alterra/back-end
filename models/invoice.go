package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	Date    string `validate:"required" json:"date" form:"date"`
	Price   string `validate:"required" json:"price" form:"price"`
	Payment string `validate:"required" json:"payment" form:"payment"`
	Type    string `validate:"required" json:"type" form:"type"`
	Status  string `validate:"required" json:"status" form:"status"`
	UserID  int    `json:"user_id" form:"user_id"`
	User    User
}

func (req *Invoice) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}

type InvoiceResponse struct {
	Date    string `validate:"required" json:"date" form:"date"`
	Price   string `validate:"required" json:"price" form:"price"`
	Payment string `validate:"required" json:"payment" form:"payment"`
	Type    string `validate:"required" json:"type" form:"type"`
	Status  string `validate:"required" json:"status" form:"status"`
	UserID  int    `json:"user_id" form:"user_id"`
}

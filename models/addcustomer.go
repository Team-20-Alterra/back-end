package models

import (
	"gorm.io/gorm"
)

type AddCustomer struct {
	gorm.Model
	BusinnesID 	int `validate:"required" json:"businnes_id" form:"businnes_id"`
	UserID  int    `json:"user_id" form:"user_id"`
	Businnes    Business
	User    User
}
type AddCustomerResponseFK struct {
	ID int `json:"id"`
	BusinnesID 	int `validate:"required" json:"businnes_id" form:"businnes_id"`
	UserID  int    `json:"user_id" form:"user_id"`
	Businnes    BusinessResponseFK `json:"businnes" form:"businnes"`
	User    UserResponseFK `json:"customer" form:"customer"`
}


type AddCustomerResponse struct {
	BusinnesID 	int `json:"businnes_id" form:"businnes_id"`
	UserID  int    `validate:"required" json:"user_id" form:"user_id"`
}
type IdCustomerResponse struct {
	BusinnesID 	int `json:"businnes_id" form:"businnes_id"`
}
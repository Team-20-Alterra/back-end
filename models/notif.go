package models

import (
	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	Title       string `validate:"required" json:"title" form:"title"`
	Body      	string `gorm:"type:text" validate:"required" json:"price" form:"price"`
	Is_readAdmin    	bool 	`gorm:"type:bool;default:false"  json:"is_readAmin" form:"is_readAmin"`
	Is_readUser    	bool 	`gorm:"type:bool;default:false"  json:"is_readUser" form:"is_readUser"`
	Status     	string `validate:"required" json:"status" form:"status"`
	InvoiceID 	uint `validate:"required" json:"customer_id" form:"customer_id"`
	Invoice     Invoice
}
type NotificationInput struct {
	gorm.Model
	Title       string `validate:"required" json:"title" form:"title"`
	Body      	string `gorm:"type:text" validate:"required" json:"price" form:"price"`
	Is_readAdmin    	bool 	`gorm:"type:bool;default:false" json:"is_readAmin" form:"is_readAmin"`
	Is_readUser    	bool 	`gorm:"type:bool;default:false" json:"is_readUser" form:"is_readUser"`
	InvoiceID 	uint `validate:"required" json:"customer_id" form:"customer_id"`
}

type NotifResponse struct {
	Is_readAdmin    	bool 	`gorm:"type:bool;default:false" json:"is_readAmin" form:"is_readAmin"`
	Is_readUser    	bool 	`gorm:"type:bool;default:false" json:"is_readUser" form:"is_readUser"`
}
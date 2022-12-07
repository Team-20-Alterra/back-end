package models

import (
	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	Title       string `validate:"required" json:"title" form:"title"`
	Body      	string `gorm:"type:text" validate:"required" json:"price" form:"price"`
	Is_read    	bool 	`validate:"required" json:"is_read" form:"is_read"`
	Type       	string `validate:"required" json:"type" form:"type"`
	Status     	string `validate:"required" json:"status" form:"status"`
	InvoiceID 	uint `validate:"required" json:"customer_id" form:"customer_id"`
	Invoice     Invoice
}

type NotifResponse struct {
	Is_read    	bool 	`validate:"required" json:"is_read" form:"is_read"`
}
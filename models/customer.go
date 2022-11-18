package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name          string    `valid:"Required;MaxSize(60)" json:"name"`
	Date_of_birth time.Time `valid:"Required" json:"date"`
	Gender        string    `valid:"Required;MaxSize(10)" json:"gender"`
	Phone         string    `valid:"Required;MaxSize(15)" json:"phone"`
	Address       string    `valid:"Required" json:"address"`
	User_id       int       `valid:"Required" json:"user_id"`
	User          User
}

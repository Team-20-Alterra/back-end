package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name          string    `valid:"Required;MaxSize(50)" json:"name"`
	Date_of_birth time.Time `valid:"Required" json:"date"`
	Email         string    `valid:"Required;MaxSize(150)" json:"email" gorm:"unique"`
	Gender        string    `valid:"Required" json:"gender"`
	Phone         string    `valid:"Required" json:"phone"`
	Address       string    `valid:"Required" json:"address"`
	Photo         string    `json:"photo"`
	Username      string    `valid:"Required;Range(6, 16)" json:"username" gorm:"unique"`
	Password      string    `valid:"Required" json:"password"`
	Role          string    `valid:"Required" json:"role"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

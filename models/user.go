package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name          string `validate:"required" json:"name" form:"name"`
	Date_of_birth string `validate:"required" json:"date" form:"date"`
	Email         string `validate:"required,email" json:"email" form:"email" gorm:"unique"`
	Gender        string `validate:"required" json:"gender" form:"gender"`
	Phone         string `validate:"required" json:"phone" form:"phone"`
	Address       string `validate:"required" json:"address" form:"address"`
	Photo         string `json:"photo" form:"photo"`
	Username      string `validate:"required" json:"username" form:"username" gorm:"unique"`
	Password      string `validate:"required" json:"password" form:"password"`
	Role          string `validate:"required" json:"role" form:"role"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

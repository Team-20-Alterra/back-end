package models

import (
	"time"

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
	Username      string `validate:"required" json:"username" form:"username"`
	Password      string `validate:"required" json:"password" form:"password"`
	Role          string `validate:"required" json:"role" form:"role"`
	PasswordResetToken string
	PasswordResetAt    time.Time
}

type UserRegister struct {
	// ID       	  int    `json:"id"`
	Name          string `validate:"required" json:"name" form:"name"`
	Email         string `validate:"required,email" json:"email" form:"email" gorm:"unique"`
	Phone         string `validate:"required" json:"phone" form:"phone"`
	Password      string `validate:"required" json:"password" form:"password"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

// ? ForgotPasswordInput struct
type ForgotPasswordInput struct {
	Email string  `validate:"required,email" json:"email" form:"email"`
}

// 👈 ResetPasswordInput struct
type ResetPasswordInput struct {
	Password        string `validate:"required" json:"password" form:"password"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required"`
}
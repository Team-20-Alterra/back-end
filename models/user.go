package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name               string `validate:"required" json:"name" form:"name"`
	Email              string `validate:"required,email" json:"email" form:"email" gorm:"unique"`
	Phone              string `validate:"required" json:"phone" form:"phone"`
	Address            string `json:"address" form:"address"`
	Photo              string `json:"photo" form:"photo"`
	Password           string `validate:"required" json:"password" form:"password"`
	Role               string `json:"role" form:"role"`
	PasswordResetToken string
	PasswordResetAt    time.Time
}

type GoogleAccount struct{
	Email         string `json:"email"`
	ID            string `json:"id"`
	Picture       string `json:"picture"`
	VerifiedEmail bool   `json:"verified_email"`
}
type UserResponseFK struct {
	ID                 int    `json:"id"`
	Name               string `validate:"required" json:"name" form:"name"`
	Email              string `validate:"required,email" json:"email" form:"email" gorm:"unique"`
	Phone              string `validate:"required" json:"phone" form:"phone"`
	Address            string `json:"address" form:"address"`
	Photo              string `json:"photo" form:"photo"`
}
type UserPreload struct {
	Name               string `validate:"required" json:"name" form:"name"`
	Email              string `validate:"required,email" json:"email" form:"email" gorm:"unique"`
	Phone              string `validate:"required" json:"phone" form:"phone"`
	Address            string `json:"address" form:"address"`
	Photo              string `json:"photo" form:"photo"`
}

type UserRegister struct {
	Name     string `validate:"required" json:"name" form:"name"`
	Email    string `validate:"required,email" json:"email" form:"email" gorm:"unique"`
	Phone    string `validate:"required" json:"phone" form:"phone"`
	Password string `validate:"required" json:"password" form:"password"`
}
type UserAdminRegister struct {
	Name     string `validate:"required" json:"name" form:"name"`
	Email    string `validate:"required,email" json:"email" form:"email" gorm:"unique"`
	Password string `validate:"required" json:"password" form:"password"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Email string `json:"email"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

// ? ForgotPasswordInput struct
type ForgotPasswordInput struct {
	Email string `validate:"required,email" json:"email" form:"email"`
}

// 👈 ResetPasswordInput struct
type ResetPasswordInput struct {
	Password        string `validate:"required" json:"password" form:"password"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required"`
}

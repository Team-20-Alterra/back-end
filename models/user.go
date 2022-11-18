package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name          string    	  `validate:"required" json:"name"`
	Date_of_birth time.Time 	  `validate:"required" json:"date"`
	Email         string    	  `validate:"required,email" json:"email" gorm:"unique"`
	Gender        string    	  `validate:"required" json:"gender"`
	Phone         string    	  `validate:"required" json:"phone"`
	Address       string    	  `validate:"required" json:"address"`
	Photo         string    	  `json:"photo"`
	Username      string    	  `validate:"required" json:"username" gorm:"unique"`
	Password      string    	  `validate:"required" json:"password"`
	Role          string    	  `validate:"required" json:"role"`
}

func (req *User) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

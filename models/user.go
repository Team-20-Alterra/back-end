package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Company_name string `valid:"Required;MaxSize(50)" json:"company_name"`
	Email        string `valid:"Required;MaxSize(150)" json:"email"`
	Address      string `valid:"Required" json:"address"`
	Username     string `valid:"Required;Range(6, 16)" json:"username"`
	Password     string `valid:"Required" json:"password"`
	Role         string `valid:"Required" json:"role"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

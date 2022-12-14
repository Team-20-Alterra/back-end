package config

import (
	"fmt"
	"geinterra/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	InitDB()
	InitialMigration()
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {
	config := Config{
		// DB_Username: "admin",
		// DB_Password: "admin123",
		// DB_Port:     "3306",
		// DB_Host:     "geinterra.cwgkf8ctyhpm.ap-northeast-1.rds.amazonaws.com",
		// DB_Name:     "geinterra_apps",
		DB_Username: "root",
		DB_Password: "",
		DB_Port:     "3306",
		DB_Host:     "localhost",
		DB_Name:     "geinterra_apps",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(
		&models.User{},
		&models.Invoice{},
		&models.Notification{},
		&models.Business{},
		&models.Bank{},
		&models.Item{},
		&models.AddCustomer{},
		&models.ListBank{},
		&models.PaymentMethod{})
}

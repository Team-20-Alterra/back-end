package main

import (
	"geinterra/config"
	mid "geinterra/middleware"
	"geinterra/routes"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
    validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
    return cv.validator.Struct(i)
}

func main() {
	config.InitDB()
	e := routes.New()
	// e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	mid.LogMiddleware(e)
	e.Logger.Fatal(e.Start(":8000"))
}

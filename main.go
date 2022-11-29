package main

import (
	"geinterra/config"
	mid "geinterra/middleware"
	"geinterra/routes"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4/middleware"
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
	e.Use(middleware.CORS())
	e.Validator = &CustomValidator{validator: validator.New()}
	mid.LogMiddleware(e)
	e.Logger.Fatal(e.Start(":8000"))
}

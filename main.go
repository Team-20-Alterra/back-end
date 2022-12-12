package main

import (
	"fmt"
	"geinterra/Seeder"
	"geinterra/config"
	mid "geinterra/middleware"
	"geinterra/routes"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
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
	Seeder.Load(config.DB)
	e := routes.New()
	e.Use(middleware.CORS())
	e.GET("/", handleMain)
	// e.GET("/login", handleGoogleLogin)
	// e.GET("/auth/:provider/callback", handleGoogleCallback)

	e.Validator = &CustomValidator{validator: validator.New()}
	mid.LogMiddleware(e)
	e.Logger.Fatal(e.Start(":8000"))
}

func handleMain(c echo.Context) error {
	var htmlIndex = `<html>
<body>
	<a href="/api/v1/login/google">Google Log In</a>
</body>
</html>`

	fmt.Fprintf(c.Response().Writer, htmlIndex)

	return c.JSON(200, "Oke")
}

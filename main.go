package main

import (
	"geinterra/config"
	mid "geinterra/middleware"
	"geinterra/routes"
)

// type CustomValidator struct {
//     validator *validator.Validate
//   }

// func (cv *CustomValidator) Validate(i interface{}) error {
//   if err := cv.validator.Struct(i); err != nil {
//     // Optionally, you could return the error to give each route more control over the status code
//     return echo.NewHTTPError(http.StatusBadRequest, err.Error())
//   }
//   return nil
// }

func main() {
	config.InitDB()
	e := routes.New()
	// e.Validator = &CustomValidator{validator: validator.New()}
	mid.LogMiddleware(e)
	e.Logger.Fatal(e.Start(":8000"))
}

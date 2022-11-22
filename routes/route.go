package routes

import (
	"geinterra/constants"
	"geinterra/controller"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	auth := e.Group("api/v1")
	auth.POST("/register", controller.CreateUserController)
	auth.POST("/login", controller.LoginController)
	
	eUser := e.Group("api/v1/users", mid.JWT([]byte(constants.SECRET_KEY)))

	// eUser.GET("", controller.GetUserController)
	eUser.GET("/profile", controller.GetUserController)
	eUser.DELETE("/profile", controller.DeleteUserController)
	eUser.PUT("/profile", controller.UpdateUserController)

	return e
}

package routes

import (
	"geinterra/controller"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	auth := e.Group("api/v1")
	auth.POST("/register", controller.CreateUserController)
	auth.POST("/login", controller.LoginController)
	
	eUser := e.Group("api/v1/users")
	// eUser.Use(mid.JWT([]byte(constants.SECRET_KEY)))

	eUser.GET("", controller.GetUserController)
	eUser.GET("/:id", controller.GetUserController)
	eUser.DELETE("/:id", controller.DeleteUserController)
	eUser.PUT("/:id", controller.UpdateUserController)

	return e
}

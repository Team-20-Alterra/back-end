package routes

import (
	"geinterra/controller"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	
	eUser := e.Group("users")
	// eUser.Use(mid.JWT([]byte(constants.SECRET_KEY)))

	eUser.GET("", controller.GetUsersController)
	eUser.GET("/:id", controller.GetUserController)
	eUser.DELETE("/:id", controller.DeleteUserController)
	eUser.PUT("/:id", controller.UpdateUserController)

	e.POST("/register", controller.CreateUserController)
	e.POST("/login", controller.LoginController)

	return e
}

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
	auth.POST("/register/admin", controller.RegisterAdminController)
	auth.POST("/register/user", controller.RegisterUserController)
	auth.POST("/login", controller.LoginController)
	auth.GET("/forgot-password", controller.ForgotPasswordController)
	// auth.GET()
	
	routes := e.Group("api/v1", mid.JWT([]byte(constants.SECRET_KEY)))

	routes.GET("/users/profile", controller.GetUserController)
	routes.DELETE("/users/profile", controller.DeleteUserController)
	routes.PUT("/users/profile", controller.UpdateUserController)

	routes.GET("/notif", controller.GetNotifController)
	routes.GET("/notif/user", controller.GetNotifByUserController)
	routes.GET("/notif/count", controller.CountNotifController)

	return e
}

package routes

import (
	"geinterra/constants"
	"geinterra/controller"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func UserRoute(e *echo.Group) {
	eUser := e.Group("users")

	eUser.Use(mid.JWT([]byte(constants.SECRET_KEY)))

	eUser.GET("/profile", controller.GetUserController)
	eUser.DELETE("/profile", controller.DeleteUserController)
	eUser.PUT("/profile", controller.UpdateUserController)
}

func InvoiceRoute(e *echo.Group) {
	eInvoice := e.Group("invoices")

	eInvoice.Use(mid.JWT([]byte(constants.SECRET_KEY)))

	eInvoice.GET("", controller.GetInvoicesController)
	eInvoice.POST("", controller.CreateInvoiceController)
	eInvoice.GET("/:id", controller.GetInvoiceController)
	eInvoice.DELETE("/:id", controller.DeleteInvoiceController)
	eInvoice.PUT("/:id", controller.UpdateInvoiceController)
}

func New() *echo.Echo {
	e := echo.New()

	v1 := e.Group("/api/v1/")
	UserRoute(v1)
	InvoiceRoute(v1)

	v1.POST("register", controller.CreateUserController)
	v1.POST("login", controller.LoginController)

	return e
}

package routes

import (
	"geinterra/controller"

	"github.com/labstack/echo"
)

func UserRoute(e *echo.Group) {
	eUser := e.Group("users")

	eUser.GET("", controller.GetUsersController)
	eUser.GET("/:id", controller.GetUserController)
	eUser.DELETE("/:id", controller.DeleteUserController)
	eUser.PUT("/:id", controller.UpdateUserController)
}

func InvoiceRoute(e *echo.Group) {
	eInvoice := e.Group("invoices")
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
	// eUser.Use(mid.JWT([]byte(constants.SECRET_KEY)))

	return e
}

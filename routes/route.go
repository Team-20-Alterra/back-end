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

	// users
	eUser.GET("", controller.GetAllUserController)
	eUser.GET("/:id", controller.GetUserByIdController)
	eUser.DELETE("/:id", controller.DeleteUserByIdController)
	// profile users
	eUser.GET("/profile", controller.GetProfileController)
	eUser.DELETE("/profile", controller.DeleteUserProfileController)
	eUser.PUT("/profile", controller.UpdateUserController)
}

func UserRole(e *echo.Group){
	roleUser := e.Group("role")

	roleUser.Use(mid.JWT([]byte(constants.SECRET_KEY)))
	roleUser.GET("/user", controller.GetUserRoleUserController)
	roleUser.GET("/admin", controller.GetUserRoleAdminController)
}

func NotifRoute(e *echo.Group){
	notif := e.Group("notif")

	notif.Use(mid.JWT([]byte(constants.SECRET_KEY)))
	notif.GET("", controller.GetNotifController)
	notif.GET("/user", controller.GetNotifByUserController)
	notif.GET("/count", controller.CountNotifController)
}

func InvoiceRoute(e *echo.Group) {
	eInvoice := e.Group("invoices")

	eInvoice.Use(mid.JWT([]byte(constants.SECRET_KEY)))

	eInvoice.GET("/coba", controller.CobaGetAll)

	eInvoice.GET("", controller.GetInvoicesController)
	eInvoice.POST("", controller.CreateInvoiceController)
	eInvoice.GET("/:id", controller.GetInvoiceController)
	eInvoice.DELETE("/:id", controller.DeleteInvoiceController)
	eInvoice.PUT("/:id", controller.UpdateInvoiceController)

	eInvoice.GET("/status/berhasil", controller.GetStatusBerhasilInvoice)
	eInvoice.GET("/status/konfirmasi", controller.GetStatusKonfirInvoice)
	eInvoice.GET("/status", controller.GetAllStatusInvoice)
	eInvoice.PUT("/update-status-bayar/:id", controller.UpdateStatusPembayaranInvoice)
	eInvoice.PUT("/update-status/:id", controller.UpdateStatusInvoice)
}

func New() *echo.Echo {
	e := echo.New()

	v1 := e.Group("/api/v1/")
	UserRoute(v1)
	InvoiceRoute(v1)
	NotifRoute(v1)
	UserRole(v1)

	v1.POST("register/admin", controller.RegisterAdminController)
	v1.POST("register/user", controller.RegisterUserController)
	v1.POST("login", controller.LoginController)
	v1.POST("forgot-password", controller.ForgotPasswordController)
	v1.PATCH("reset-password/:resetToken", controller.ResetPassword)
	return e
}

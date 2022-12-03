package routes

import (
	"geinterra/constants"
	"geinterra/controller"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func BankRoute(e *echo.Group) {
	eBank := e.Group("banks")

	eBank.Use(mid.JWT([]byte(constants.SECRET_KEY)))

	eBank.GET("", controller.GetBanksController)
	eBank.POST("", controller.CreateBankController)
	eBank.GET("/:id", controller.GetBankController)
	eBank.DELETE("/:id", controller.DeleteBankController)
	eBank.PUT("/:id", controller.UpdateBankController)
}

func BusinessRoute(e *echo.Group) {
	eBusiness := e.Group("business")

	eBusiness.Use(mid.JWT([]byte(constants.SECRET_KEY)))

	eBusiness.GET("", controller.GetBusinesssController)
	eBusiness.POST("", controller.CreateBusinessController)
	eBusiness.GET("/:id", controller.GetBusinessController)
	eBusiness.DELETE("/:id", controller.DeleteBusinessController)
	eBusiness.PUT("/:id", controller.UpdateBusinessController)
}

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

func UserRole(e *echo.Group) {
	roleUser := e.Group("role")

	roleUser.Use(mid.JWT([]byte(constants.SECRET_KEY)))

	roleUser.GET("/user", controller.GetUserRoleUserController)
	roleUser.GET("/admin", controller.GetUserRoleAdminController)
}

func NotifRoute(e *echo.Group) {
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

func ItemRoute(e *echo.Group){
	eItem := e.Group("item")
	eItem.Use(mid.JWT([]byte(constants.SECRET_KEY)))
	eItem.GET("", controller.GetItemController)
	eItem.GET("/invoice/:id", controller.GetItemByInvoiceController)
	eItem.POST("", controller.CreateItemController)
	eItem.DELETE("/:id", controller.DeleteItemController)
	eItem.PUT("/:id", controller.UpdateItemController)
}

func AddCustomerRoute(e *echo.Group){
	eCustomer := e.Group("add-customer")

	eCustomer.Use(mid.JWT([]byte(constants.SECRET_KEY)))

	eCustomer.GET("/businness", controller.GetCustomerByBusinness)
	eCustomer.POST("", controller.AddCustomerController)
	eCustomer.DELETE("/:id", controller.DeleteCustomer)

}

func New() *echo.Echo {
	e := echo.New()

	v1 := e.Group("/api/v1/")
	UserRoute(v1)
	InvoiceRoute(v1)
	NotifRoute(v1)
	UserRole(v1)
	BusinessRoute(v1)
	BankRoute(v1)
	ItemRoute(v1)
	AddCustomerRoute(v1)

	v1.POST("register/admin", controller.RegisterAdminController)
	v1.POST("register/user", controller.RegisterUserController)
	v1.POST("login/admin", controller.LoginAdminController)
	v1.POST("login", controller.LoginController)
	v1.POST("forgot-password", controller.ForgotPasswordController)
	v1.PATCH("reset-password/:resetToken", controller.ResetPassword)
	return e
}

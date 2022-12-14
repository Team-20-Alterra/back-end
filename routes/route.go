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

	// eBank.GET("", controller.GetBanksController)
	eBank.POST("", controller.CreateBankController)
	eBank.GET("/:id", controller.GetBankController)
	eBank.DELETE("/:id", controller.DeleteBankController)
	eBank.PUT("/:id", controller.UpdateBankController)
}

func BusinessRoute(e *echo.Group) {
	eBusiness := e.Group("business")

	eBusiness.Use(mid.JWT([]byte(constants.SECRET_KEY)))

	eBusiness.GET("", controller.GetBusinesssController)
	eBusiness.GET("/user", controller.GetBusinessByUserController)
	eBusiness.GET("/:id", controller.GetBusinessController)

	eBusiness.DELETE("", controller.DeleteBusinessController)
	eBusiness.PUT("", controller.UpdateBusinessController)

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
	notif.GET("/busines", controller.GetNotifByAdminController)
	notif.GET("/user/:id", controller.GetNotifByIdUser)
	notif.GET("/admin/:id", controller.GetNotifByIdAdmin)
	notif.GET("/count-user", controller.CountNotifUserController)
	notif.GET("/count-admin", controller.CountNotifAdminController)
	notif.DELETE("/:id", controller.DeleteNotifController)
}

func InvoiceRoute(e *echo.Group) {
	eInvoice := e.Group("invoices")

	eInvoice.Use(mid.JWT([]byte(constants.SECRET_KEY)))

	eInvoice.GET("", controller.GetInvoicesController)
	eInvoice.POST("", controller.CreateInvoiceController)
	eInvoice.GET("/:id", controller.GetInvoiceController)
	eInvoice.DELETE("/:id", controller.DeleteInvoiceController)
	eInvoice.PUT("/:id", controller.UpdateInvoiceController)

	// get status for admin
	eInvoice.GET("/status", controller.GetAllStatusAdminInvoice)
	eInvoice.GET("/status/berhasil", controller.GetStatusBerhasilInvoice)
	eInvoice.GET("/status/on-proses", controller.GetStatusOnProsesInvoice)
	eInvoice.GET("/status/pending", controller.GetStatusPendingInvoice)
	eInvoice.GET("/status/gagal", controller.GetStatusGagalInvoice)
	eInvoice.GET("/status/konfir", controller.GetStatusMenungguKonfirInvoice)
	// get status for customer
	eInvoice.GET("/status/customer", controller.GetAllStatusCustomerInvoice)
	eInvoice.GET("/status/berhasil/customer", controller.GetStatusBerhasilInvoiceCustomer)
	eInvoice.GET("/status/on-proses/customer", controller.GetStatusOnProsesInvoiceCustomer)
	eInvoice.GET("/status/pending/customer", controller.GetStatusPendingInvoiceCustomer)
	eInvoice.GET("/status/gagal/customer", controller.GetStatusGagalInvoiceCustomer)

	eInvoice.PUT("/update-status-bayar/:id", controller.UpdateStatusPembayaranInvoice)
	eInvoice.PUT("/update-status/:id", controller.UpdateStatusInvoice)

	// seacrh
	eInvoice.GET("/search", controller.SearchInvoice)
	eInvoice.GET("/search/status/customer", controller.SearchInvoiceStatusForCustomer)
	eInvoice.GET("/search/status/admin", controller.SearchInvoiceStatusForAdmin)

	// count
	eInvoice.GET("/count", controller.GetCountSubtotalAll)
	eInvoice.GET("/count/berhasil", controller.GetCountSubtotalBerhasil)
	eInvoice.GET("/count/gagal", controller.GetCountSubtotalGagal)

}

func ItemRoute(e *echo.Group) {
	eItem := e.Group("item")
	eItem.Use(mid.JWT([]byte(constants.SECRET_KEY)))
	eItem.GET("", controller.GetItemController)
	eItem.GET("/invoice/:id", controller.GetItemByInvoiceController)
	eItem.POST("", controller.CreateItemController)
	eItem.DELETE("/:id", controller.DeleteItemController)
	eItem.PUT("/:id", controller.UpdateItemController)
}

func AddCustomerRoute(e *echo.Group) {
	eCustomer := e.Group("add-customer")

	eCustomer.Use(mid.JWT([]byte(constants.SECRET_KEY)))

	eCustomer.GET("/businness", controller.GetCustomerByBusinness)
	eCustomer.POST("", controller.AddCustomerController)
	eCustomer.DELETE("/:id", controller.DeleteCustomer)
}

func ListBank(e *echo.Group) {
	eListBank := e.Group("list-bank")

	eListBank.Use(mid.JWT([]byte(constants.SECRET_KEY)))

	eListBank.GET("", controller.GetListBanksController)
	eListBank.GET("/:id", controller.GetListBankByIdController)
	eListBank.GET("/:business", controller.GetListBankBusinessController)
	eListBank.GET("/businness", controller.GetListBankByBusinessController)
	eListBank.POST("", controller.CreateListBankController)
}

func Checkout(e *echo.Group){
	eCheckout := e.Group("checkout")

	eCheckout.Use(mid.JWT([]byte(constants.SECRET_KEY)))

	eCheckout.POST("", controller.CreateCheckoutController)
	eCheckout.PUT("/:id", controller.UpdateCheckoutController)
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
	ListBank(v1)
	Checkout(v1)

	v1.GET("login/google", controller.LoginGoogleController)
	v1.POST("register/admin", controller.RegisterAdminController)
	v1.POST("register/user", controller.RegisterUserController)
	v1.POST("login/admin", controller.LoginAdminController)
	v1.POST("login", controller.LoginController)
	v1.POST("forgot-password", controller.ForgotPasswordController)
	v1.PATCH("reset-password", controller.ResetPassword)
	v1.POST("register/busines", controller.RegisterBusinessController)
	v1.GET("banks", controller.GetBanksController)

	e.GET("/auth/:provider/callback", controller.HandleGoogleCallbackController)
	return e
}

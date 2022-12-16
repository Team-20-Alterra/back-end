package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"geinterra/config"
	"geinterra/gomail"
	"geinterra/models"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func GetInvoicesController(c echo.Context) error {
	var invoices []models.Invoice

	// user := c.Get("user").(*jwt.Token)
	// claims := user.Claims.(jwt.MapClaims)

	// id, _ := claims["id"]

	if err := config.DB.Preload("Businnes.User").Preload("User").Preload("Item").Preload("Checkout.ListBank.Bank").Find(&invoices).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success get all invoices",
		"invoices": invoices,
	})
}

// get invoice by id
func GetInvoiceController(c echo.Context) error {
	var invoice models.Invoice

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Where("id = ?", id).Preload("Businnes.User").Preload("User").Preload("Item").Preload("Checkout.ListBank.Bank").First(&invoice).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  false,
			"message": "Record not found!",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "success get Invoice",
		"data":    invoice,
	})
}

// get status by busines
func GetStatusBerhasilInvoice(c echo.Context) error {
	var invoice []models.Invoice
	var busines models.Business

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

	// cek already busines
	if err := config.DB.Where("user_id = ?", id).First(&busines).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Business already exist",
			"data":    nil,
		})
	}

	status := "Berhasil"

	if err := config.DB.Where("businnes_id = ?", busines.ID).Where("status = ?", status).Preload("Businnes.User").Preload("User").Preload("Item").Preload("Checkout.ListBank.Bank").Find(&invoice).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  false,
			"message": "Record not found!",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "success get Invoice by status berhasil",
		"data":    invoice,
	})
}
func GetStatusMenungguKonfirInvoice(c echo.Context) error {
	var invoice []models.Invoice
	var busines models.Business

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

	// cek already busines
	if err := config.DB.Where("user_id = ?", id).First(&busines).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Business already exist",
			"data":    nil,
		})
	}

	status := "Menunggu Konfirmasi"

	if err := config.DB.Where("businnes_id = ?", busines.ID).Where("status = ?", status).Preload("Businnes.User").Preload("User").Preload("Item").Preload("Checkout.ListBank.Bank").Find(&invoice).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  false,
			"message": "Record not found!",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "success get Invoice by status berhasil",
		"data":    invoice,
	})
}

func GetStatusOnProsesInvoice(c echo.Context) error {
	var invoice []models.Invoice
	var busines models.Business

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

	// cek already busines
	if err := config.DB.Where("user_id = ?", id).First(&busines).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Business already exist",
			"data":    nil,
		})
	}

	fmt.Println(busines.ID)

	status := "On Proses"

	if err := config.DB.Where("businnes_id = ?", busines.ID).Where("status = ?", status).Preload("Businnes.User").Preload("User").Preload("Item").Preload("Checkout.ListBank.Bank").Find(&invoice).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  false,
			"message": "Record not found!",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "success get Invoice by status on proses",
		"data":    invoice,
	})
}

func GetStatusPendingInvoice(c echo.Context) error {
	var invoice []models.Invoice
	var busines models.Business

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

	// cek already busines
	if err := config.DB.Where("user_id = ?", id).First(&busines).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Business already exist",
			"data":    nil,
		})
	}

	fmt.Println(busines.ID)

	status := "Pending"

	if err := config.DB.Where("businnes_id = ?", busines.ID).Where("status = ?", status).Preload("Businnes.User").Preload("User").Preload("Item").Preload("Checkout.ListBank.Bank").Find(&invoice).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  false,
			"message": "Record not found!",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "success get Invoice by status pending",
		"data":    invoice,
	})
}

func GetStatusGagalInvoice(c echo.Context) error {
	var invoice []models.Invoice
	var busines models.Business

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

	// cek already busines
	if err := config.DB.Where("user_id = ?", id).First(&busines).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Business already exist",
			"data":    nil,
		})
	}

	status := "Gagal"

	if err := config.DB.Where("businnes_id = ?", busines.ID).Where("status = ?", status).Preload("Businnes.User").Preload("User").Preload("Item").Preload("Checkout.ListBank.Bank").Find(&invoice).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  false,
			"message": "Record not found!",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "success get Invoice by status pending",
		"data":    invoice,
	})
}

func GetAllStatusAdminInvoice(c echo.Context) error {
	var invoice []models.Invoice
	var busines models.Business

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

	// cek already busines
	if err := config.DB.Where("user_id = ?", id).First(&busines).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Business already exist",
			"data":    nil,
		})
	}

	if err := config.DB.Where("businnes_id = ?", busines.ID).Preload("Businnes.User").Preload("User").Preload("Item").Preload("Checkout.ListBank.Bank").Find(&invoice).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  false,
			"message": "Record not found!",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "success get Invoice by status",
		"data":    invoice,
	})
}

// count subtotal
func GetCountSubtotalBerhasil(c echo.Context) error {
	var busines models.Business

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

	// cek already busines
	if err := config.DB.Where("user_id = ?", id).First(&busines).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Business already exist",
			"data":    nil,
		})
	}

	status := "Berhasil"

	var result int64
	row := config.DB.Table("invoices").Select("sum(subtotal)").Where("businnes_id = ?", busines.ID).Where("status = ?", status).Row()
	row.Scan(&result)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "success get sub total by status berhasil",
		"data":    result,
	})
}
func GetCountSubtotalGagal(c echo.Context) error {
	var busines models.Business

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

	// cek already busines
	if err := config.DB.Where("user_id = ?", id).First(&busines).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Business already exist",
			"data":    nil,
		})
	}

	status := "Gagal"

	var result int64
	row := config.DB.Table("invoices").Select("sum(subtotal)").Where("businnes_id = ?", busines.ID).Where("status = ?", status).Row()
	row.Scan(&result)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "success get sub total by status gagal",
		"data":    result,
	})
}
func GetCountSubtotalAll(c echo.Context) error {
	var busines models.Business

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

	// cek already busines
	if err := config.DB.Where("user_id = ?", id).First(&busines).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Business already exist",
			"data":    nil,
		})
	}

	var result int64
	row := config.DB.Table("invoices").Select("sum(subtotal)").Where("businnes_id = ?", busines.ID).Row()
	row.Scan(&result)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "success get all count sub total",
		"data":    result,
	})
}

// get status by customer
func GetAllStatusCustomerInvoice(c echo.Context) error {
	var invoice []models.Invoice

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

	if err := config.DB.Where("user_id = ?", id).Preload("Businnes.User").Preload("User").Preload("Item").Preload("Checkout.ListBank.Bank").Find(&invoice).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  false,
			"message": "Record not found!",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "success get Invoice by status",
		"data":    invoice,
	})
}

func GetStatusBerhasilInvoiceCustomer(c echo.Context) error {
	var invoice []models.Invoice

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

	status := "Berhasil"

	if err := config.DB.Where("user_id = ?", id).Where("status = ?", status).Preload("Businnes.User").Preload("User").Preload("Item").Preload("Checkout.ListBank.Bank").Find(&invoice).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  false,
			"message": "Record not found!",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "success get Invoice by status berhasil",
		"data":    invoice,
	})
}

func GetStatusOnProsesInvoiceCustomer(c echo.Context) error {
	var invoice []models.Invoice

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

	status := "On Proses"

	if err := config.DB.Where("user_id = ?", id).Where("status = ?", status).Preload("Businnes.User").Preload("User").Preload("Item").Preload("Checkout.ListBank.Bank").Find(&invoice).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  false,
			"message": "Record not found!",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "success get Invoice by status on proses",
		"data":    invoice,
	})
}

func GetStatusPendingInvoiceCustomer(c echo.Context) error {
	var invoice []models.Invoice

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

	status := "Pending"

	if err := config.DB.Where("user_id = ?", id).Where("status = ?", status).Preload("Businnes.User").Preload("User").Preload("Item").Preload("Checkout.ListBank.Bank").Find(&invoice).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  false,
			"message": "Record not found!",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "success get Invoice by status pending",
		"data":    invoice,
	})
}

func GetStatusGagalInvoiceCustomer(c echo.Context) error {
	var invoice []models.Invoice

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

	status := "Gagal"

	if err := config.DB.Where("user_id = ?", id).Where("status = ?", status).Preload("Businnes.User").Preload("User").Preload("Item").Preload("Checkout.ListBank.Bank").Find(&invoice).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  false,
			"message": "Record not found!",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "success get Invoice by status gagal",
		"data":    invoice,
	})
}

// create invoice
func CreateInvoiceController(c echo.Context) error {
	var busines models.Business
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"].(float64)

	var invoice models.InvoiceResponse
	c.Bind(&invoice)

	// cek already busines
	if err := config.DB.Where("user_id = ?", id).First(&busines).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  false,
			"message": "Business not found",
			"data":    nil,
		})
	}

	invoice.BusinnesID = int(busines.ID)

	invoiceReal := models.Invoice{Payment: invoice.Payment, Type: invoice.Type, Note: invoice.Note, Status: invoice.Status, UserID: int(id), BusinnesID: invoice.BusinnesID}

	if err := config.DB.Create(&invoiceReal).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":    true,
		"message":   "success create new invoice",
		"IdInvoice": invoiceReal.ID,
	})
}

// update invoice
func UpdateInvoiceController(c echo.Context) error {
	var invoice models.Invoice

	id, _ := strconv.Atoi(c.Param("id"))

	var input models.InvoiceUpdate

	body, _ := ioutil.ReadAll(c.Request().Body)
	err := json.Unmarshal(body, &input)
	if err != nil {
		return err
	}

	if err := config.DB.Model(&invoice).Where("id = ?", id).First(&invoice).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Invoice not found!",
			"data":    nil,
		})
	}

	input.Status = "Menunggu Konfirmasi"

	invoiceReal := models.Invoice{Total: input.Total, Discount: input.Discount, Note: input.Note, Subtotal: input.Subtotal, Status: input.Status, UserID: input.UserID}

	if err := config.DB.Model(&invoice).Where("id = ?", id).Updates(&invoiceReal).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Invoice not found!",
			"data":    nil,
		})
	}

	var users models.User

	// cari user
	if err := config.DB.Where("id = ?", input.UserID).First(&users).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "User not found!",
			"data":    nil,
		})
	}

	// TO DO Create Notif user
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}

	emailTo := users.Email

	fmt.Println(emailTo)

	data := struct {
		ReceiverName string
	}{
		ReceiverName: users.Name,
	}

	gomail.OAuthGmailService()
	status, err := gomail.SendEmailOAUTH2(emailTo, data, "notif.html")
	if err != nil {
		log.Println(err)
	}
	if status {
		log.Println("Email sent successfully using OAUTH")
	}

	// create notif
	var inputNotif models.NotificationInput

	err = json.Unmarshal(body, &inputNotif)
	if err != nil {
		return err
	}

	fmt.Println(invoice)
	fmt.Println(invoice.UserID)

	inputNotif.InvoiceID = uint(id)

	notif := models.Notification{Title: inputNotif.Title, Body: inputNotif.Body, InvoiceID: inputNotif.InvoiceID}

	fmt.Println(notif)

	if err := c.Validate(inputNotif); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Create(&notif).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "update success",
	})
}

// update pembayaran invoice
func UpdateStatusPembayaranInvoice(c echo.Context) error {
	var invoice models.Invoice

	id, _ := strconv.Atoi(c.Param("id"))

	// cek invoice
	if err := config.DB.Where("id = ?", id).First(&invoice).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  false,
			"message": "Invoice not found!",
			"data":    nil,
		})
	}

	var input models.InvoicePembayaranStatus
	c.Bind(&input)

	fileHeader, _ := c.FormFile("payment")
	if fileHeader != nil {
		file, _ := fileHeader.Open()

		ctx := context.Background()

		cldService, _ := cloudinary.NewFromURL(os.Getenv("URL_CLOUDINARY"))

		resp, _ := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})

		input.Payment = resp.SecureURL
	}

	// input.DatePay = time.Now().String()

	// invoiceUpdate := models.Invoice{Status: input.Status, Payment: input.Payment, DatePay: input.DatePay}
	// if err := c.Validate(input); err != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]any{
	// 		"status":  false,
	// 		"message": err.Error(),
	// 		"data":    nil,
	// 	})
	// }
	invoiceUpdate := models.Invoice{Status: input.Status, Payment: input.Payment}
	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": err.Error(),
			"data":    nil,
		})
	}
	if err := config.DB.Model(&invoice).Where("id = ?", id).Updates(&invoiceUpdate).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Record not found!",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "update success",
	})
}

func UpdateStatusInvoice(c echo.Context) error {
	var invoice models.Invoice

	id, _ := strconv.Atoi(c.Param("id"))

	var input models.InvoiceStatus
	c.Bind(&input)

	// invoiceUpdate := models.Invoice{StatusInvoice: input.StatusInvoice, Status: input.Status}
	// if err := c.Validate(input); err != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]any{
	// 		"status":  false,
	// 		"message": err.Error(),
	// 		"data":    nil,
	// 	})
	// }
	invoiceUpdate := models.Invoice{Status: input.Status}
	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": err.Error(),
			"data":    nil,
		})
	}
	// cek invoice
	if err := config.DB.Where("id = ?", id).First(&invoice).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  false,
			"message": "Invoice not found!",
			"data":    nil,
		})
	}

	if err := config.DB.Model(&invoice).Where("id = ?", id).Updates(&invoiceUpdate).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Record not found!",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "update success",
	})
}

func DeleteInvoiceController(c echo.Context) error {
	var Invoices models.Invoice

	id, _ := strconv.Atoi(c.Param("id"))

	// cek invoice
	if err := config.DB.Where("id = ?", id).First(&Invoices).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  false,
			"message": "Invoice not found!",
			"data":    nil,
		})
	}
	if err := config.DB.Delete(&Invoices, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Record not found!",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "success delete invoice",
	})
}

// seacrh
func SearchInvoice(c echo.Context) error {
	var invoices []models.Invoice
	var busines models.Business

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

	// cek already busines
	if err := config.DB.Where("user_id = ?", id).First(&busines).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Business already exist",
			"data":    nil,
		})
	}

	searctData := c.QueryParam("search")
	search := c.QueryParam("data")

	if err := config.DB.Where("businnes_id = ?", busines.ID).Where("type LIKE ? OR no_invoice LIKE ? OR created_at LIKE ? OR price between ? AND ?", "%"+searctData+"%", "%"+searctData+"%", "%"+searctData+"%", searctData, search).Preload("Businnes").Preload("User").Preload("Item").Preload("Checkout.ListBank.Bank").Find(&invoices).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  false,
			"message": "Record not found!",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "success get all invoices",
		"data":    invoices,
	})
}
func SearchInvoiceStatusForCustomer(c echo.Context) error {
	var invoices []models.Invoice

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

	searctData := c.QueryParam("status")

	if err := config.DB.Where("user_id = ?", id).Where("status LIKE ? ", "%"+searctData+"%").Preload("Businnes").Preload("User").Preload("Item").Preload("Checkout.ListBank.Bank").Find(&invoices).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  false,
			"message": "Record not found!",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "success get all invoices",
		"data":    invoices,
	})
}
func SearchInvoiceStatusForAdmin(c echo.Context) error {
	var invoices []models.Invoice
	var busines models.Business

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

	// cek already busines
	if err := config.DB.Where("user_id = ?", id).First(&busines).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Business already exist",
			"data":    nil,
		})
	}

	searctData := c.QueryParam("status")

	if err := config.DB.Where("businnes_id = ?", busines.ID).Where("status LIKE ? ", "%"+searctData+"%").Preload("Businnes").Preload("User").Preload("Item").Preload("Checkout.ListBank.Bank").Find(&invoices).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  false,
			"message": "Record not found!",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "success get all invoices",
		"data":    invoices,
	})
}

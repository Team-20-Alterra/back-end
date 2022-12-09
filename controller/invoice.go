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
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func GetInvoicesController(c echo.Context) error {
	var invoices []models.Invoice

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

	if err := config.DB.Where("user_id = ?", id).Preload("User").Preload("Item").Find(&invoices).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success get all invoices",
		"invoices": invoices,
	})
}

func GetInvoiceController(c echo.Context) error {
	var invoice models.Invoice

	id, _ := strconv.Atoi(c.Param("id"))

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	userId, _ := claims["id"]

	if err := config.DB.Where("id = ?", id).Where("user_id = ?", userId).Preload("User").Joins("Item").First(&invoice).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get Invoice",
		"invoice": invoice,
	})
}

func GetStatusKonfirInvoice(c echo.Context) error {
	var invoice []models.Invoice

	status := "Menunggu Konfirmasi"

	if err := config.DB.Where("status = ?", status).Preload("User").First(&invoice).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get Invoice by status",
		"invoice": invoice,
	})
}

func GetStatusDiprosesInvoice(c echo.Context) error {
	var invoice []models.Invoice

	status := "Diproses"

	if err := config.DB.Where("status = ?", status).Preload("User").First(&invoice).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get Invoice by status",
		"invoice": invoice,
	})
}

func GetStatusPendingInvoice(c echo.Context) error {
	var invoice []models.Invoice

	status := "Pending"

	if err := config.DB.Where("status = ?", status).Preload("User").Preload("Item").First(&invoice).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get Invoice by status",
		"invoice": invoice,
	})
}

func GetAllStatusInvoice(c echo.Context) error {
	var invoice []models.Invoice

	konfir := "Menunggu Konfirmasi"
	gagal := "Gagal"
	proses := "Dalam Proses"
	jatuh := "Jatuh Tempo"

	if err := config.DB.Where("status = ?", konfir).Or("status = ?", gagal).Or("status = ?", proses).Or("status = ?", jatuh).Preload("User").First(&invoice).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get Invoice by status",
		"invoice": invoice,
	})
}

func GetStatusBerhasilInvoice(c echo.Context) error {
	var invoice []models.Invoice

	status := "Berhasil"

	if err := config.DB.Where("status = ?", status).Preload("User").First(&invoice).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get Invoice by status berhasil",
		"invoice": invoice,
	})
}

func CreateInvoiceController(c echo.Context) error {
	var busines models.Business
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"].(float64)

	var invoice models.InvoiceResponse
	c.Bind(&invoice)

	// cek already busines
	if err := config.DB.Where("user_id = ?", id).First(&busines).Error; err != nil {
		return c.JSON(http.StatusAlreadyReported, map[string]any{
			"status":  false,
			"message": "Business already exist",
			"data":    nil,
		})
	}

	now := time.Now()

	// day := busines.Due_Date

	// billing add 10 day
	toAdd := 240 * time.Hour

	newTime := now.Add(toAdd)

	invoice.BillingDate = newTime.String()

	invoice.BusinnesID = int(busines.ID)

	invoiceReal := models.Invoice{DatePay: invoice.DatePay, Price: invoice.Price, Payment: invoice.Payment, Type: invoice.Type, Status: invoice.Status, UserID: int(id), BusinnesID: invoice.BusinnesID, BillingDate: invoice.BillingDate}

	if err := config.DB.Create(&invoiceReal).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "success create new invoice",
		"data":    invoiceReal,
	})
}

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

	invoiceReal := models.Invoice{Price:input.Price, Payment: input.Payment,Type: input.Type,Status: input.Status, UserID: input.UserID, NoInvoice: input.NoInvoice}

	if err := config.DB.Model(&invoice).Where("id = ?", id).Updates(&invoiceReal).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Invoice not found!",
			"data":    nil,
		})
	}

	// if err := config.DB.Where("id = ?", id).Where("user_id = ?", idUser).First(&invoice).Error; err != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]any{
	// 		"status":  false,
	// 		"message": "Record not found!",
	// 		"data":    nil,
	// 	})
	// }

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

func DeleteInvoiceController(c echo.Context) error {
	var Invoices models.Invoice

	id, _ := strconv.Atoi(c.Param("id"))

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

func UpdateStatusPembayaranInvoice(c echo.Context) error {
	var invoice models.Invoice

	id, _ := strconv.Atoi(c.Param("id"))

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

	invoiceUpdate := models.Invoice{StatusInvoice: input.StatusInvoice}
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

func FilterByDataController(c echo.Context) error {
	var invoices []models.Invoice

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

	if err := config.DB.Where("user_id = ?", id).Find(&invoices).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	if err := config.DB.Order("date").Find(&invoices).Error; err != nil {
		panic("failed to retrieve data")
	}

	var dataInv map[string]interface{}

	for key, value := range invoices {
		dataInv[string(key)] = value
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success get all invoices",
		"invoices": invoices,
	})
}

func FilterByDate(c echo.Context) error {
	var invoices []models.Invoice

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

	date := c.Param("date")

	if err := config.DB.Where("date = ?", date).Where("user_id = ?", id).Order("date").Find(&invoices).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	// if err := config.DB.Order("date").Find(&invoices).Error; err != nil {
	// 	panic("failed to retrieve data")
	// }

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success get all invoices",
		"invoices": invoices,
	})
}

func FilterByStatus(c echo.Context) error {
	var invoices []models.Invoice

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

	status := c.Param("status")

	if err := config.DB.Where("status = ?", status).Where("user_id = ?", id).Order("date").Find(&invoices).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success get all invoices",
		"invoices": invoices,
	})
}

func FilterByPrice(c echo.Context) error {
	var invoices []models.Invoice

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

	var input map[string]interface{}

	body, _ := ioutil.ReadAll(c.Request().Body)
	err := json.Unmarshal(body, &input)
	if err != nil {
		log.Print(err)
		return nil
	}

	log.Print(time.Now().Month())
	log.Print(time.Now().Day())
	log.Print(time.Now().Year())
	log.Print(time.Now().Clock())

	if input["price_min"] == 0 {
		if err := config.DB.Where("price < ?", input["price_max"]).Where("user_id = ?", id).Order("price").Find(&invoices).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
		}
	} else if input["price_max"] == 0 {
		if err := config.DB.Where("price > ?", input["price_min"]).Where("user_id = ?", id).Order("price").Find(&invoices).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
		}
	} else {
		if err := config.DB.Where("? < price < ?", input["price_min"], input["price_max"]).Where("user_id = ?", id).Order("price").Find(&invoices).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success get all invoices",
		"invoices": invoices,
	})
}

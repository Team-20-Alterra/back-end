package controller

import (
	"context"
	"encoding/json"
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

func CreateInvoiceController(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"].(float64)

	var invoice models.InvoiceResponse
	c.Bind(&invoice)

	// fileHeader, _ := c.FormFile("payment")
	// if fileHeader != nil {
	// 	file, _ := fileHeader.Open()

	// 	ctx := context.Background()

	// 	cldService, _ := cloudinary.NewFromURL(os.Getenv("URL_CLOUDINARY"))

	// 	resp, _ := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})

	// 	invoice.Payment = resp.SecureURL
	// }

	// date := "2006-01-02"
	// dob, _ := time.Parse(date, invoice.Date)

	// invoice.Date = dob.String()
	// invoice.Status = "Menunggu Konfirmasi"

	// if err := c.Validate(invoice); err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	// }

	invoiceReal := models.Invoice{Date: invoice.Date, Price: invoice.Price, Payment: invoice.Payment, Type: invoice.Type, Status: invoice.Status, UserID: int(id)}

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

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	idUser, _ := claims["id"]

	id, _ := strconv.Atoi(c.Param("id"))

	var input models.Invoice
	c.Bind(&input)

	// fileHeader, _ := c.FormFile("payment")
	// if fileHeader != nil {
	// 	file, _ := fileHeader.Open()

	// 	ctx := context.Background()

	// 	cldService, _ := cloudinary.NewFromURL(os.Getenv("URL_CLOUDINARY"))

	// 	resp, _ := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})

	// 	input.Payment = resp.SecureURL
	// }

	// date := "12-12-2001"
	// log.Print(input.Date)
	// dob, _ := time.Parse(date, input.Date.(string))

	// input.Date = dob
	// log.Print(dob)
	input.Status = "Menunggu Konfirmasi"

	if err := config.DB.Model(&invoice).Where("id = ?", id).Where("user_id = ?", idUser).Updates(&input).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Record not found!",
			"data":    nil,
		})
	}

	if err := config.DB.Where("id = ?", id).Where("user_id = ?", idUser).First(&invoice).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Record not found!",
			"data":    nil,
		})
	}

	var users models.User

	if err := config.DB.Where("id = ?", invoice.UserID).First(&users).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Record not found!",
			"data":    nil,
		})
	}
	// TO DO Create Notif
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}

	emailTo := users.Email

	data := struct {
		ReceiverName string
	}{
		ReceiverName: input.User.Name,
	}

	gomail.OAuthGmailService()
	status, err := gomail.SendEmailOAUTH2(emailTo, data, "notif.html")
	if err != nil {
		log.Println(err)
	}
	if status {
		log.Println("Email sent successfully using OAUTH")
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

func CobaGetAll(c echo.Context) error {
	var invoice []models.Invoice

	if err := config.DB.Order("year, month").Find(&invoice).Error; err != nil {
		panic("failed to retrieve data")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "update success",
		"data":    invoice,
	})
}

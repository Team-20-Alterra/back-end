package controller

import (
	"context"
	"geinterra/config"
	"geinterra/models"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/labstack/echo/v4"
)

func GetCheckoutController(c echo.Context) error {
	var checkouts []models.Checkout

	if err := config.DB.Preload("Businnes").Preload("User").Preload("Item").Find(&checkouts).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":    "success get all checkouts",
		"checkouts": checkouts,
	})
}

func CreateCheckoutController(c echo.Context) error {
	var checkout models.CheckoutInput
	var invoice models.Invoice
	var list models.ListBank
	var check models.Checkout
	c.Bind(&checkout)

	// check listbank
	if err := config.DB.Where("id = ?", checkout.ListBankID).First(&list).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "ListBank not found!",
			"data":    nil,
		})
	}

	// check invoice
	if err := config.DB.Where("id = ?", checkout.InvoiceID).First(&invoice).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Invoice not found!",
			"data":    nil,
		})
	}

	// check checkout
	if err := config.DB.Where("invoice_id = ?", checkout.InvoiceID).First(&check).Error; err == nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "This invoice already has checkout",
			"data":    nil,
		})
	}

	now := time.Now()
	// billing due date
	toAdd := 24 * time.Hour

	newTime := now.Add(toAdd)

	checkout.BillingDate = newTime.String()

	checkoutReal := models.Checkout{ListBankID: checkout.ListBankID, BillingDate: checkout.BillingDate, InvoiceID: checkout.InvoiceID}

	if err := config.DB.Create(&checkoutReal).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "success create new checkout",
		"data":    checkout,
	})
}

func UpdateCheckoutController(c echo.Context) error {
	// update checkout
	var check models.Checkout
	var checkout models.CheckoutUpdate
	var invoice models.Invoice
	var input models.InvoicePembayaranStatus

	c.Bind(&input)
	c.Bind(&checkout)
	id, _ := strconv.Atoi(c.Param("id"))
	// cari checkout
	if err := config.DB.Where("id = ?", id).First(&check).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Record not found!",
			"data":    nil,
		})
	}

	// update checkout
	checkout.DatePay = time.Now().String()

	if err := c.Validate(checkout); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	checkoutUpdate := models.Checkout{DatePay: checkout.DatePay}
	if err := config.DB.Model(&check).Where("id = ?", id).Updates(&checkoutUpdate).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Record not found!",
			"data":    nil,
		})
	}

	// update invoice
	fileHeader, _ := c.FormFile("payment")
	if fileHeader != nil {
		file, _ := fileHeader.Open()

		ctx := context.Background()

		cldService, _ := cloudinary.NewFromURL(os.Getenv("URL_CLOUDINARY"))

		resp, _ := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})

		input.Payment = resp.SecureURL
	}

	invoiceUpdate := models.Invoice{Status: input.Status, Payment: input.Payment, Type: input.Type}
	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": err.Error(),
			"data":    nil,
		})
	}
	if err := config.DB.Model(&invoice).Where("id = ?", check.InvoiceID).Updates(&invoiceUpdate).Error; err != nil {
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

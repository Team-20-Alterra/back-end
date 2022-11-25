package controller

import (
	"context"
	"geinterra/config"
	"geinterra/models"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/labstack/echo"
)

func GetInvoicesController(c echo.Context) error {
	var invoices []models.Invoice

	if err := config.DB.Find(&invoices).Error; err != nil {
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

	if err := config.DB.Where("id = ?", id).First(&invoice).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get Invoice",
		"invoice": invoice,
	})
}

func CreateInvoiceController(c echo.Context) error {
	sortResponse := []string{"status", "message", "data"}
	sort.Strings(sortResponse)

	var invoice models.Invoice
	c.Bind(&invoice)

	fileHeader, _ := c.FormFile("payment")
	log.Println(fileHeader.Filename)

	file, _ := fileHeader.Open()

	ctx := context.Background()

	cldService, _ := cloudinary.NewFromURL(os.Getenv("URL_CLOUDINARY"))

	resp, _ := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})
	log.Println(resp.SecureURL)

	invoice.Payment = resp.SecureURL

	date := "2006-01-02"
	dob, _ := time.Parse(date, invoice.Date)

	invoice.Date = dob.String()
	invoice.CreatedAt = time.Now()
	invoice.UpdatedAt = time.Now()

	// if err := c.Validate(invoice); err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	// }

	if err := config.DB.Model(&invoice).Create(&invoice).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Create failed!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		sortResponse[0]: true,
		sortResponse[1]: "success create new invoice",
		sortResponse[2]: invoice,
	})
}

func UpdateInvoiceController(c echo.Context) error {
	var invoice models.Invoice

	id, _ := strconv.Atoi(c.Param("id"))

	var input models.Invoice
	c.Bind(&input)

	fileHeader, _ := c.FormFile("payment")
	log.Println(fileHeader.Filename)

	file, _ := fileHeader.Open()

	ctx := context.Background()

	cldService, _ := cloudinary.NewFromURL(os.Getenv("URL_CLOUDINARY"))

	resp, _ := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})
	log.Println(resp.SecureURL)

	input.Payment = resp.SecureURL

	if err := config.DB.Model(&invoice).Where("id = ?", id).Updates(input).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "update success",
	})
}

func DeleteInvoiceController(c echo.Context) error {
	var Invoices models.Invoice

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Delete(&Invoices, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete Invoice",
	})
}

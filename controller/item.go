package controller

import (
	"geinterra/config"
	"geinterra/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetItemController(c echo.Context) error {
	var item []models.Item

	if err := config.DB.Find(&item).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status": true,
		"message": "success get all item by invoice",
		"data": item,
	})
}
func GetItemByInvoiceController(c echo.Context) error {
	var item []models.Item

	invoiceId, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Where("invoice_id = ?", invoiceId).Preload("Invoice.User").Find(&item).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status": true,
		"message": "success get all item by invoice",
		"data": item,
	})
}

func CreateItemController(c echo.Context) error {
	var item models.ItemResponse
	var invoice models.Invoice
	
	id := item.InvoiceID

	if err := config.DB.Where("id = ?", id).First(&invoice).Error; err != nil {
		return c.JSON(http.StatusAlreadyReported, map[string] any {
			"status": false,
			"message": "Failed invoice",
			"data": nil,
		})
	}

	c.Bind(&item)

	if err := c.Validate(item); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	createItem := models.Item{Name: item.Name, Amount: item.Amount, UnitPrice: item.UnitPrice, TotalPrice: item.TotalPrice, InvoiceID: item.InvoiceID}

	if err := config.DB.Create(&createItem).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "success create new item",
		"data":    item,
	})
}

func UpdateItemController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var input models.ItemResponse

	item := models.Item{
		Name: input.Name,
		Amount: input.Amount,
		UnitPrice: input.UnitPrice,
		TotalPrice: input.TotalPrice,
		InvoiceID: input.InvoiceID,
	}

	if err := config.DB.Model(&input).Where("id = ?", id).Updates(item).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status": false,
			"message": "Record not found!",
			"data": nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"message": "update success",
	})
}

func DeleteItemController(c echo.Context) error {
	var item []models.Item

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Where("id = ?", id).First(&item).Delete(&item).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]any {
			"status": false,
			"message": "Record not found!",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"message": "success delete item",
	})
}
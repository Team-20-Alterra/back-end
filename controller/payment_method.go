package controller

import (
	"geinterra/config"
	"geinterra/models"
	"net/http"
	"sort"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetPaymentMethodsController(c echo.Context) error {
	var payment []models.PaymentMethod

	if err := config.DB.Find(&payment).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success get all payment",
		"payment": payment,
	})
}

func GetPaymentMethodController(c echo.Context) error {
	var payment models.PaymentMethod

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Where("id = ?", id).First(&payment).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get payment",
		"payment": payment,
	})
}

func CreatePaymentMethodController(c echo.Context) error {
	var payment models.PaymentMethod
	c.Bind(&payment)

	if err := config.DB.Create(&payment).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "success create new payment",
		"data":    payment,
	})
}

func UpdatePaymentMethodController(c echo.Context) error {
	sortResponse := []string{"status", "message", "data"}
	sort.Strings(sortResponse)

	id, _ := strconv.Atoi(c.Param("id"))

	var payment models.PaymentMethod

	var input models.PaymentMethod
	c.Bind(&input)

	if err := config.DB.Model(&payment).Where("id = ?", id).Updates(input).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			sortResponse[0]: false,
			sortResponse[1]: "Record not found!",
			sortResponse[2]: nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		sortResponse[0]: true,
		sortResponse[1]: "update success",
	})
}

func DeletePaymentMethodController(c echo.Context) error {
	sortResponse := []string{"status", "message", "data"}
	sort.Strings(sortResponse)

	var payment models.PaymentMethod

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Delete(&payment, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			sortResponse[0]: false,
			sortResponse[1]: "Record not found!",
			sortResponse[2]: nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		sortResponse[0]: true,
		sortResponse[1]: "success delete payment",
	})
}

func GetPaymentMethodByBankID(c echo.Context) error {
	sortResponse := []string{"status", "message", "data"}
	sort.Strings(sortResponse)

	var payment models.PaymentMethod
	c.Bind(&payment)

	if err := config.DB.Where("bank_id = ?", payment.BankID).Preload("Bank").First(&payment).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			sortResponse[0]: false,
			sortResponse[1]: "Record not found!",
			sortResponse[2]: nil,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		sortResponse[0]: true,
		sortResponse[1]: "success get payment method",
		sortResponse[2]: payment,
	})
}

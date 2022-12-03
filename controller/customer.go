package controller

import (
	"encoding/json"
	"fmt"
	"geinterra/config"
	"geinterra/models"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddCustomerController(c echo.Context) error {

	var customer models.AddCustomer
	var addCustomer models.AddCustomerResponse
	var business models.Business

	body, _ := ioutil.ReadAll(c.Request().Body)
	err := json.Unmarshal(body, &addCustomer)
	if err != nil {
		return err
	}

	id := addCustomer.BusinnesID

	if err := config.DB.Where("id = ?", id).First(&business).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]any {
			"status": false,
			"message": "Record not found!",
		})
	}

	newCustomer := models.AddCustomer{
		UserID: addCustomer.UserID,
		BusinnesID: int(business.ID),
	}

	if err := config.DB.Where("user_id = ? AND businnes_id = ?", addCustomer.UserID, addCustomer.BusinnesID).First(&customer).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]any {
			"status": false,
			"message": "Record not found!",
		})
	}

	if err := c.Validate(addCustomer); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	if err := config.DB.Model(&customer).Create(&newCustomer).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Create failed!",
			"data":    nil,
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  true,
		"message": "success create add new customer",
		"data":    addCustomer,
	})
}

func GetCustomerByBusinness(c echo.Context) error {
	var customer []models.AddCustomer
	var cusId models.IdCustomerResponse
	c.Bind(&cusId)

	fmt.Println(cusId.BusinnesID)

	if err := config.DB.Where("businnes_id = ?", cusId.BusinnesID).Preload("Businnes.Bank").Preload("User").Find(&customer).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status": false,
			"message": "Record not found!",
			"data": nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status": true,
		"message": "success get all customer by business",
		"data": customer,
	})
}

func DeleteCustomer(c echo.Context) error {
	var customer []models.AddCustomer

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Where("id = ?", id).First(&customer).Delete(&customer).Error; err != nil {
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
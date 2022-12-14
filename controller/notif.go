package controller

import (
	"fmt"
	"geinterra/config"
	"geinterra/models"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetNotifController(c echo.Context) error {
	var notif []models.Notification

	if err := config.DB.Preload("Invoice.User").Find(&notif).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status": true,
		"message": "success get all notif",
		"data": notif,
	})
}

// get notif by user
func GetNotifByUserController(c echo.Context) error {
	// var invoice models.Invoice
	var notif []models.Notification

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

	if err := config.DB.Joins("Invoice").Where("Invoice.user_id = ?", id).Preload("Invoice.User").Find(&notif).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string] any {
			"status": false,
			"message": "Record not found!" ,
			"data": nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status": true,
		"message": "success get all notif by user",
		"data": notif,
	})
}

func GetNotifByAdminController(c echo.Context) error {
	var invoice models.Invoice
	var notif []models.Notification
	var busines models.Business

	
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	
	id, _ := claims["id"]

	fmt.Println(busines.UserID)

	// cek busines
	if err := config.DB.Where("user_id = ?", id).First(&busines).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Business Not Found",
			"data":    nil,
		})
	}

	if err := config.DB.Joins("JOIN notifications on notifications.invoice_id=invoices.id").
		Where("invoices.businnes_id=?", busines.ID).
		Group("invoices.id").Preload("User").Find(&invoice).Error; err != nil {
			return c.JSON(http.StatusNotFound, map[string]any{
				"status": false,
				"message": err.Error(),
				"data": nil,
			})
	}
	
	// validate not found invoice
	if invoice.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status": false,
			"message": "Invoice record not found!",
			"data": nil,
		})		
	}

	if err := config.DB.Joins("Invoice").Where("Invoice.businnes_id = ?", busines.ID).Preload("Invoice.User").Find(&notif).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string] any {
			"status": false,
			"message": "Record not found!" ,
			"data": nil,
		})
	}

	fmt.Println(busines.ID)

	return c.JSON(http.StatusOK, map[string]any{
		"status": true,
		"message": "success get all notif by user",
		"data": notif,
	})
}

// get id notif by user
func GetNotifByIdUser(c echo.Context) error {
	var notif models.Notification
	// var input models.NotifResponseUser

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Where("id = ?", id).Preload("Invoice.User").First(&notif).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status": false,
			"message": "Record not found!",
			"data": nil,
		})
	}

	notifUpdate := models.Notification{
		Is_readUser: true,
	}

	if err := config.DB.Model(&notif).Where("id = ?", id).Updates(notifUpdate).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status": false,
			"message": "Record not found!",
			"data": nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"message": "success get notif",
		"data":    notif,
	})
}
func GetNotifByIdAdmin(c echo.Context) error {
	var notif models.Notification
	// var input models.NotifResponseUser

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Where("id = ?", id).Preload("Invoice.User").First(&notif).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status": false,
			"message": "Record not found!",
			"data": nil,
		})
	}

	notifUpdate := models.Notification{
		Is_readAdmin: true,
	}

	if err := config.DB.Model(&notif).Where("id = ?", id).Updates(notifUpdate).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status": false,
			"message": "Record not found!",
			"data": nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"message": "success get notif",
		"data":    notif,
	})
}

// count notif user
func CountNotifUserController(c echo.Context) error {
	var notif []models.Notification
	
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	
	id, _ := claims["id"]

	var count int64

	if err := config.DB.Joins("Invoice").Where("Invoice.user_id = ?", id).Where("is_read_user = ?", false).Find(&notif).Count(&count).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string] any {
			"status": false,
			"message": "Record not found!" ,
			"data": nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status": true,
		"message": "success get count notif by user",
		"data": count,
	})
}
func CountNotifAdminController(c echo.Context) error {
	// var invoice models.Invoice
	var notif []models.Notification
	var busines models.Business

	
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	
	id, _ := claims["id"]

	fmt.Println(busines.UserID)

	// cek busines
	if err := config.DB.Where("user_id = ?", id).First(&busines).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Business Not Found",
			"data":    nil,
		})
	}

	var count int64

	if err := config.DB.Joins("Invoice").Where("Invoice.businnes_id  = ?", busines.UserID).Where("is_read_admin = ?", false).Preload("Invoice.User").Find(&notif).Count(&count).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string] any {
			"status": false,
			"message": "Record not found!" ,
			"data": nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status": true,
		"message": "success get count notif by user",
		"data": count,
	})
}

func DeleteNotifController(c echo.Context) error {	
	var notif models.Notification
	
	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Where("id = ?", id).First(&notif).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any {
			"status": false,
			"message": "Notif not found!",
			"data": nil,
		})
	}

	if err := config.DB.Delete(&notif, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any {
			"status": false,
			"message": "Record not found!",
			"data": nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"message": "success delete notif",
	})
}
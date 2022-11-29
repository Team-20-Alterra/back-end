package controller

import (
	"geinterra/config"
	"geinterra/models"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetNotifController(c echo.Context) error {
	var notif []models.Notification

	if err := config.DB.Find(&notif).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status": true,
		"message": "success get all notif",
		"data": notif,
	})
}
func GetNotifByUserController(c echo.Context) error {
	var notif models.Notification

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

	if err := config.DB.Where("id = ?", id).First(&notif).Error; err != nil {
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
func CountNotifController(c echo.Context) error {
	var notif models.Notification

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

	var count int64

	if err := config.DB.Where("id = ?", id).First(&notif).Count(&count).Error; err != nil {
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
	
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

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
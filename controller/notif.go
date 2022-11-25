package controller

import (
	"geinterra/config"
	"geinterra/models"
	"net/http"
	"sort"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetNotifController(c echo.Context) error {
	sortResponse := []string{"status", "message", "data"}
	sort.Strings(sortResponse)

	var notif []models.Notification

	if err := config.DB.Find(&notif).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]any{
		sortResponse[0]: true,
		sortResponse[1]: "success get all notif",
		sortResponse[2]: notif,
	})
}
func GetNotifByUserController(c echo.Context) error {
	sortResponse := []string{"status", "message", "data"}
	sort.Strings(sortResponse)

	var notif models.Notification

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

	if err := config.DB.Where("id = ?", id).First(&notif).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string] any {
			sortResponse[0]: false,
			sortResponse[1]: "Record not found!" ,
			sortResponse[2]: nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		sortResponse[0]: true,
		sortResponse[1]: "success get all notif by user",
		sortResponse[2]: notif,
	})
}
func CountNotifController(c echo.Context) error {
	sortResponse := []string{"status", "message", "data"}
	sort.Strings(sortResponse)

	var notif models.Notification

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

	var count int64

	if err := config.DB.Where("id = ?", id).First(&notif).Count(&count).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string] any {
			sortResponse[0]: false,
			sortResponse[1]: "Record not found!" ,
			sortResponse[2]: nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		sortResponse[0]: true,
		sortResponse[1]: "success get count notif by user",
		sortResponse[2]: count,
	})
}
func DeleteNotifController(c echo.Context) error {
	sortResponse := []string{"status", "message", "data"}
	sort.Strings(sortResponse)
	var notif models.Notification

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

	if err := config.DB.Delete(&notif, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any {
			sortResponse[0]: false,
			sortResponse[1]: "Record not found!",
			sortResponse[2]: nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		sortResponse[0]: true,
		sortResponse[1]: "success delete notif",
	})
}



package controller

import (
	"geinterra/config"
	"geinterra/middleware"
	"geinterra/models"
	"net/http"

	"github.com/labstack/echo"
)

func LoginController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	err := config.DB.Where("username = ? AND password = ?", user.Username, user.Password).First(&user).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	token, err := middleware.CreateToken(int(user.ID), user.Username, user.Email, user.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	userResponse := models.UserResponse{int(user.ID), user.Username, user.Email, user.Role, token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "Berhasil Login",
		"data":    userResponse,
	})
}

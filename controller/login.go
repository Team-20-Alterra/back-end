package controller

import (
	"geinterra/config"
	"geinterra/middleware"
	"geinterra/models"
	"net/http"
	"sort"

	"github.com/labstack/echo"
)

func LoginController(c echo.Context) error {
	sortResponse := []string{"status","message", "data"}
	sort.Strings(sortResponse)
	
	user := models.User{}
	c.Bind(&user)

	err := config.DB.Where("username = ? AND password = ?", user.Username, user.Password).First(&user).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			sortResponse[0]: false,
			sortResponse[1]: err.Error(),
			sortResponse[2]: nil,
		})
	}



	token, err := middleware.CreateToken(int(user.ID), user.Username, user.Email, user.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			sortResponse[0]: false,
			sortResponse[1]: err.Error(),
			sortResponse[2]: nil,
		})
	}

	userResponse := models.UserResponse{int(user.ID), user.Username, user.Email, user.Role, token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		sortResponse[0]: true,
		sortResponse[1]: "Berhasil Login",
		sortResponse[2]: userResponse,
	})
}

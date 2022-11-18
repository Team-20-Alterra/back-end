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
		var helper models.HelperResponse
		helper.Status = false
		helper.Message = err.Error()
		// helper.Data = nil
		return c.JSON(http.StatusInternalServerError, helper)
	}

	token, err := middleware.CreateToken(int(user.ID), user.Username, user.Email, user.Role)
	if err != nil {
		var helper models.HelperResponse
		helper.Status = false
		helper.Message = err.Error()
		// helper.Data = nil
		return c.JSON(http.StatusInternalServerError, helper)
	}

	userResponse := models.UserResponse{int(user.ID), user.Username, user.Email, user.Role, token}
	var helper models.HelperResponse
	helper.Status = true
	helper.Message = err.Error()
	helper.Data = userResponse

	return c.JSON(http.StatusOK, helper)
}

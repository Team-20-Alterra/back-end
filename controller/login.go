package controller

import (
	"encoding/json"
	"geinterra/config"
	"geinterra/middleware"
	"geinterra/models"
	"io/ioutil"
	"net/http"
	"sort"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func LoginController(c echo.Context) error {
	sortResponse := []string{"status", "message", "data"}
	sort.Strings(sortResponse)

	var input models.User
	body, _ := ioutil.ReadAll(c.Request().Body)
	error := json.Unmarshal(body, &input)
	if error != nil {
		return error
	}

	user := models.User{}

	err := config.DB.Where("username = ?", input.Username).First(&user).Error

	match := CheckPasswordHash(input.Password, user.Password)

	err = config.DB.Where("username = ? AND ?", user.Username, match).First(&user).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			sortResponse[0]: false,
			sortResponse[1]: err.Error(),
			sortResponse[2]: nil,
		})
	}

	token, err := middleware.CreateToken(int(user.ID), user.Username, user.Email, user.Role)
	// token, err := middleware.
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

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

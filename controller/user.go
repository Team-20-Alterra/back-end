package controller

import (
	"encoding/json"
	"geinterra/config"
	"geinterra/models"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func GetUsersController(c echo.Context) error {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success get all users",
		"users":  users,
	})
}

func GetUserController(c echo.Context) error {
	var users models.User

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Where("id = ?", id).First(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"users":   users,
	})
}

func CreateUserController(c echo.Context) error {
	var user models.User
	var customer models.Customer
	var input map[string]interface{}

	body, _ := ioutil.ReadAll(c.Request().Body)
	err := json.Unmarshal(body, &input)
	if err != nil {
		log.Error("json body is empty")
		return nil
	}

	birth := input["date_of_birth"].(string)
	dateFormat := "02/01/2006"
	tgl, _ := time.Parse(dateFormat, birth)
	input["date_of_birth"] = tgl
	input["created_at"] = time.Now()
	input["updated_at"] = time.Now()

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Create failed!")
	}
	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Create failed!")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success create new user",
		"user":     user,
		"customer": customer,
	})
}

func UpdateUserController(c echo.Context) error {
	var users models.User

	id, _ := strconv.Atoi(c.Param("id"))

	var input models.User
	c.Bind(&input)

	if err := config.DB.Model(&users).Where("id = ?", id).Updates(input).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "update success",
	})
}

func DeleteUserController(c echo.Context) error {
	var users models.User

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Delete(&users, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})
}

package controller

import (
	"encoding/json"
	"geinterra/config"
	"geinterra/models"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
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
	sortResponse := []string{"status","message", "data"}
	sort.Strings(sortResponse)

	var user models.User
	body, _ := ioutil.ReadAll(c.Request().Body)
	err := json.Unmarshal(body, &user);if err != nil {
	return err
	}

	email := user.Email
	username := user.Username

	if err := config.DB.Where("email = ?", email).First(&user).Error; err == nil {
		return echo.NewHTTPError(http.StatusAlreadyReported, "Email Sudah ada")
	}	

	if err := config.DB.Where("username = ?", username).First(&user).Error; err == nil {
		return echo.NewHTTPError(http.StatusAlreadyReported, "Username Sudah ada")
	}

	//hashing password
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost )
	
	birth := user.Date_of_birth
	dateFormat := "02/01/2006"
	tgl, _ := time.Parse(dateFormat, birth.String())
	user.Password =  string(hash)
	user.Date_of_birth = tgl
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()


	
	if err := config.DB.Model(&user).Create(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Create failed!")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		sortResponse[0]: true,
		sortResponse[1]: "success create new user",
		sortResponse[2]: user,
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

func hashPassword(input string) string {
	password := []byte(input)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}
package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"geinterra/config"
	_middlewares "geinterra/middleware"
	"geinterra/models"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/labstack/echo/v4"
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

	getUser:= _middlewares.GetUser(c)

	fmt.Println(getUser)

	id := getUser.ID
	// id, _ := strconv.Atoi(c.Param("id"))
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

    date := "2006-01-02"
    dob, _ := time.Parse(date, user.Date_of_birth)

	user.Date_of_birth = dob.String()
	user.Password =  string(hash)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

    if err := c.Validate(user); err != nil {
        // return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		return c.JSON(http.StatusBadRequest, map[string]any {
			sortResponse[0]: false,
			sortResponse[1]: err.Error(),
			sortResponse[2]: nil,
		})
    }
	
	if err := config.DB.Model(&user).Create(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Create failed!")
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		sortResponse[0]: true,
		sortResponse[1]: "success create new user",
		sortResponse[2]: user,
	})
}

func UpdateUserController(c echo.Context) error {
	var users models.User

	// getUser:= c.Get("user").(*jwt.Token)

	id, _ := strconv.Atoi(c.Param("id"))

	var input models.User
	c.Bind(&input)

	fileHeader, _ := c.FormFile("photo")
	log.Println(fileHeader.Filename)

	file, _ := fileHeader.Open()

	ctx := context.Background()

	cldService, _ := cloudinary.NewFromURL(os.Getenv("URL_CLOUDINARY"))

	resp, _ := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})
	log.Println(resp.SecureURL)

	input.Photo = resp.SecureURL

	if err := config.DB.Model(&users).Where("id = ?", id).Updates(input).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "update success",
	})
}

func DeleteUserController(c echo.Context) error {
	var users models.User

	// getUser:= c.Get("user").(*jwt.Token)

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Delete(&users, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})
}

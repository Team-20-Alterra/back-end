package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"geinterra/config"
	"geinterra/models"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func GetUserController(c echo.Context) error {
	sortResponse := []string{"status", "message", "data"}
	sort.Strings(sortResponse)

	var users models.User

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

	if err := config.DB.Where("id = ?", id).First(&users).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			sortResponse[0]: false,
			sortResponse[1]: "Record not found!",
			sortResponse[2]: nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		sortResponse[0]: true,
		sortResponse[1]: "success get user",
		sortResponse[2]: users,
	})
}

func CreateUserController(c echo.Context) error {
	sortResponse := []string{"status", "message", "data"}
	sort.Strings(sortResponse)

	var user models.User
	body, _ := ioutil.ReadAll(c.Request().Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		return err
	}

	email := user.Email
	username := user.Username

	if err := config.DB.Where("email = ?", email).First(&user).Error; err == nil {
		return c.JSON(http.StatusAlreadyReported, map[string]any{
			sortResponse[0]: false,
			sortResponse[1]: "Email Sudah ada",
			sortResponse[2]: nil,
		})
	}

	if err := config.DB.Where("username = ?", username).First(&user).Error; err == nil {
		return c.JSON(http.StatusAlreadyReported, map[string]any{
			sortResponse[0]: false,
			sortResponse[1]: "Username Sudah ada",
			sortResponse[2]: nil,
		})
	}

	//hashing password
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	date := "2006-01-02"
	dob, _ := time.Parse(date, user.Date_of_birth)

	user.Date_of_birth = dob.String()
	user.Password = string(hash)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if err := c.Validate(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			sortResponse[0]: false,
			sortResponse[1]: err.Error(),
			sortResponse[2]: nil,
		})
	}

	if err := config.DB.Model(&user).Create(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			sortResponse[0]: false,
			sortResponse[1]: "Create failed!",
			sortResponse[2]: nil,
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		sortResponse[0]: true,
		sortResponse[1]: "success create new user",
		sortResponse[2]: user,
	})
}

func UpdateUserController(c echo.Context) error {
	sortResponse := []string{"status", "message", "data"}
	sort.Strings(sortResponse)
	var users models.User

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	fmt.Println("data", claims["id"])

	id, _ := claims["id"]

	var input models.User
	c.Bind(&input)

	fileHeader, _ := c.FormFile("photo")
	if fileHeader != nil {
		file, _ := fileHeader.Open()

		ctx := context.Background()

		cldService, _ := cloudinary.NewFromURL(os.Getenv("URL_CLOUDINARY"))

		resp, _ := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})

		input.Photo = resp.SecureURL

	}

	if err := config.DB.Model(&users).Where("id = ?", id).Updates(input).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			sortResponse[0]: false,
			sortResponse[1]: "Record not found!",
			sortResponse[2]: nil,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		sortResponse[0]: true,
		sortResponse[1]: "update success",
	})
}

func DeleteUserController(c echo.Context) error {
	sortResponse := []string{"status", "message", "data"}
	sort.Strings(sortResponse)
	var users models.User

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	fmt.Println("data", claims["id"])

	id, _ := claims["id"]

	if err := config.DB.Delete(&users, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			sortResponse[0]: false,
			sortResponse[1]: "Record not found!",
			sortResponse[2]: nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		sortResponse[0]: true,
		sortResponse[1]: "success delete user",
	})
}

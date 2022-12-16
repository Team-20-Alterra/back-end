package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"geinterra/config"
	"geinterra/models"
	"geinterra/utils"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func GetProfileController(c echo.Context) error {
	var users models.UserResponseFK

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

	if err := config.DB.Model(&models.User{}).Select("*").Where("id = ?", id).Scan(&users).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status": "User not found!",
			"data": nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"message": "success get user",
		"data": users,
	})
}

func CreateUserController(c echo.Context) error {
	var user models.User
	body, _ := ioutil.ReadAll(c.Request().Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		return err
	}

	email := user.Email
	// username := user.Username

	if err := config.DB.Where("email = ?", email).First(&user).Error; err == nil {
		return c.JSON(http.StatusBadRequest, map[string] any {
			"status": false,
			"message": "Email Sudah ada",
			"data": nil,
		})
	}

	// if err := config.DB.Where("username = ?", username).First(&user).Error; err == nil {
	// 	return c.JSON(http.StatusBadRequest, map[string] any {
	// 		"status": false,
	// 		"message": "Username Sudah ada",
	// 		"data": nil,
	// 	})
	// }

	//hashing password
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	// date := "2006-01-02"
	// dob, _ := time.Parse(date, user.Date_of_birth)

	// user.Date_of_birth = dob.String()
	user.Password = string(hash)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

    if err := c.Validate(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any {
			"status": false,
			"message": err.Error(),
			"data": nil,
		})
    }
	
	if err := config.DB.Model(&user).Create(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string] any {
			"status": false,
			"message": "Create failed!",
			"data": nil,
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": true,
		"message": "success create new user",
		"data": user,
	})
}

func UpdateUserController(c echo.Context) error {	
	var users models.User
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

	var input models.UserUpdate
	c.Bind(&input)

	fileHeader, _ := c.FormFile("photo")
	if fileHeader != nil {
		file, _ := fileHeader.Open()

		ctx := context.Background()

		cldService, _ := cloudinary.NewFromURL(os.Getenv("URL_CLOUDINARY"))

		resp, _ := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})

		input.Photo = resp.SecureURL

	}

	if users.Email != input.Email {
		email := input.Email
	
		if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string] any {
				"status": false,
				"message": "Email Sudah ada",
				"data": nil,
			})
		}
	}

	if users.Phone != input.Phone {
		phone := input.Phone
	
		if err := config.DB.Where("phone = ?", phone).First(&user).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string] any {
				"status": false,
				"message": "Phone Sudah ada",
				"data": nil,
			})
		}
	}

	hash, _ := utils.HashPassword(input.Password)

	input.Password = hash

	if err := config.DB.Model(&users).Where("id = ?", id).Updates(input).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status": false,
			"message": "Failed to save data",
			"data": nil,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": true,
		"message": "update success",
	})
}

func DeleteUserProfileController(c echo.Context) error {	var users models.User

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	fmt.Println("data", claims["id"])

	id, _ := claims["id"]

	if err := config.DB.Delete(&users, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status": false,
			"message": "Record not found!",
			"data": nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"message": "success delete user",
	})
}

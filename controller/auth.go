package controller

import (
	"encoding/json"
	"geinterra/config"
	"geinterra/gomail"
	"geinterra/middleware"
	"geinterra/models"
	"geinterra/utils"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/thanhpk/randstr"
)

func LoginController(c echo.Context) error {
	var input models.User
	body, _ := ioutil.ReadAll(c.Request().Body)
	error := json.Unmarshal(body, &input)
	if error != nil {
		return error
	}

	user := models.User{}

	err := config.DB.Where("email = ?", input.Email).First(&user).Error

	match := utils.CheckPasswordHash(input.Password, user.Password)

	err = config.DB.Where("email = ? AND ?", user.Email, match).First(&user).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": false,
			"message": "Incorrect Email or Password",
			"data": nil,
		})
	}

	token, err := middleware.CreateToken(int(user.ID), user.Username, user.Email, user.Role)
	// token, err := middleware.
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": false,
			"message": err.Error(),
			"data": nil,
		})
	}

	userResponse := models.UserResponse{int(user.ID), user.Username, user.Email, user.Role, token}

	return c.JSON(http.StatusOK, map[string]any{
		"status": true,
		"message": "Berhasil Login",
		"data": userResponse,
	})
}
func LoginAdminController(c echo.Context) error {
	var input models.User
	body, _ := ioutil.ReadAll(c.Request().Body)
	error := json.Unmarshal(body, &input)
	if error != nil {
		return error
	}

	user := models.User{}

	err := config.DB.Where("email = ?", input.Email).First(&user).Error

	match := utils.CheckPasswordHash(input.Password, user.Password)

	err = config.DB.Where("email = ? AND ?", user.Email, match).First(&user).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": false,
			"message": "Incorrect Email or Password",
			"data": nil,
		})
	}

	roleUser := "Admin"

	err = config.DB.Where("role = ?", roleUser).First(&user).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": false,
			"message": "Only admins can enter",
			"data": nil,
		})
	}

	token, err := middleware.CreateToken(int(user.ID), user.Username, user.Email, user.Role)
	// token, err := middleware.
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": false,
			"message": err.Error(),
			"data": nil,
		})
	}

	userResponse := models.UserResponse{int(user.ID), user.Username, user.Email, user.Role, token}

	return c.JSON(http.StatusOK, map[string]any{
		"status": true,
		"message": "Berhasil Login",
		"data": userResponse,
	})
}

func RegisterAdminController(c echo.Context) error {


	var user models.User
	var userRegister models.UserRegister

	body, _ := ioutil.ReadAll(c.Request().Body)
	err := json.Unmarshal(body, &userRegister)
	if err != nil {
		return err
	}

	email := userRegister.Email

	if err := config.DB.Where("email = ?", email).First(&user).Error; err == nil {
		return c.JSON(http.StatusAlreadyReported, map[string] any {
			"status": false,
			"message": "Email already exist",
			"data": nil,
		})
	}
	phone := userRegister.Phone

	if err := config.DB.Where("phone = ?", phone).First(&user).Error; err == nil {
		return c.JSON(http.StatusAlreadyReported, map[string] any {
			"status": false,
			"message": "Phone already exist",
			"data": nil,
		})
	}

	hash, _ := utils.HashPassword(userRegister.Password)
	
	newUser := models.User{
		Name: userRegister.Name,
		Date_of_birth: "",
		Email: userRegister.Email,
		Gender: "",
		Phone: userRegister.Phone,
		Address: "",
		Photo: "",
		Username: "",
		Password: string(hash),
		Role: "Admin",
	}

    if err := c.Validate(userRegister); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any {
			"status": false,
			"message": err.Error(),
			"data": nil,
		})
    }
	
	if err := config.DB.Model(&user).Create(&newUser).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string] any {
			"status": false,
			"message": "Create failed!",
			"data": nil,
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": true,
		"message": "success create new user",
		"data": newUser,
	})
}
func RegisterUserController(c echo.Context) error {


	var user models.User
	var userRegister models.UserRegister

	body, _ := ioutil.ReadAll(c.Request().Body)
	err := json.Unmarshal(body, &userRegister)
	if err != nil {
		return err
	}

	email := userRegister.Email

	if err := config.DB.Where("email = ?", email).First(&user).Error; err == nil {
		return c.JSON(http.StatusAlreadyReported, map[string] any {
			"status": false,
			"message": "Email already exist",
			"data": nil,
		})
	}
	phone := userRegister.Phone

	if err := config.DB.Where("phone = ?", phone).First(&user).Error; err == nil {
		return c.JSON(http.StatusAlreadyReported, map[string] any {
			"status": false,
			"message": "Phone already exist",
			"data": nil,
		})
	}

	hash, _ := utils.HashPassword(userRegister.Password)
	
	newUser := models.User{
		Name: userRegister.Name,
		Date_of_birth: "",
		Email: userRegister.Email,
		Gender: "",
		Phone: userRegister.Phone,
		Address: "",
		Photo: "",
		Username: "",
		Password: string(hash),
		Role: "User",
	}

    if err := c.Validate(userRegister); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any {
			"status": false,
			"message": err.Error(),
			"data": nil,
		})
    }
	
	if err := config.DB.Model(&user).Create(&newUser).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string] any {
			"status": false,
			"message": "Create failed!",
			"data": nil,
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": true,
		"message": "success create new user",
		"data": newUser,
	})
}

func ForgotPasswordController(c echo.Context) error {
	var users models.User

	var input models.ForgotPasswordInput
	c.Bind(&input)
	email := input.Email
    
	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any {
			"status": false,
			"message": err.Error(),
			"data": nil,
		})
    }

	if err := config.DB.Where("email = ?", email).First(&users).Error; err != nil {
		return c.JSON(http.StatusAlreadyReported, map[string] any {
			"status": false,
			"message": "Email Tidak Ditemukan",
			"data": nil,
		})
	}

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Error getting env, %v", err)
	// }

	// Generate Verification Code
	resetToken := randstr.String(20)

	passwordResetToken := utils.Encode(resetToken)
	users.PasswordResetToken = passwordResetToken
	users.PasswordResetAt = time.Now().Add(time.Minute * 15)
	config.DB.Save(&users)

	emailTo := email

		data := struct {
			ReceiverName string
			Link 		 string
		}{
			ReceiverName: users.Name,
			Link: "https://ginap-mu.vercel.app/new-password/" + resetToken,
		}

		gomail.OAuthGmailService()
		status, err := gomail.SendEmailOAUTH2(emailTo, data, "template.html")
		if err != nil {
			log.Println(err)
		}
		if status {
				log.Println("Email sent successfully using OAUTH")
		}
	return c.JSON(http.StatusOK, map[string]any{
		"status": true,
		"message": "Success, check your email right now",
		"data": nil,
	})
}

func ResetPassword(ctx echo.Context) error {
	var payload *models.ResetPasswordInput
	resetToken := ctx.Param("resetToken")

	if err := ctx.Bind(&payload); err != nil {
		
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"status": false, 
			"message": err.Error(),
		})
	} 

	if err := ctx.Validate(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any {
			"status": false,
			"message": err.Error(),
			"data": nil,
		})
    }

	if payload.Password != payload.PasswordConfirm {
		
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"status": false, 
			"message": "Passwords do not match",
		})
	}

	hashedPassword, _ := utils.HashPassword(payload.Password)

	passwordResetToken := utils.Encode(resetToken)

	var updatedUser models.User
	result := config.DB.First(&updatedUser, "password_reset_token = ? AND password_reset_at > ?", passwordResetToken, time.Now())
	if result.Error != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"status": false, 
			"message": "The reset token is invalid or has expired",
		})
	}

	updatedUser.Password = hashedPassword
	updatedUser.PasswordResetToken = ""
	config.DB.Save(&updatedUser)

	// ctx.SetCookie("token", "", -1, "/", "localhost", false, true)

	return ctx.JSON(http.StatusOK, map[string]any{
		"status": true, 
		"message": "Password data updated successfully",
	})
}




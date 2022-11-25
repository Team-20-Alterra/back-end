package controller

import (
	"encoding/json"
	"geinterra/config"
	"geinterra/gomail"
	"geinterra/middleware"
	"geinterra/models"
	"io/ioutil"
	"log"
	"net/http"
	"sort"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func LoginController(c echo.Context) error {
	sortResponse := []string{"status", "message", "data"}
	// sort.Strings(sortResponse)

	// fmt.Println(sortResponse)

	var input models.User
	body, _ := ioutil.ReadAll(c.Request().Body)
	error := json.Unmarshal(body, &input)
	if error != nil {
		return error
	}

	user := models.User{}

	err := config.DB.Where("email = ?", input.Email).First(&user).Error

	match := CheckPasswordHash(input.Password, user.Password)

	err = config.DB.Where("email = ? AND ?", user.Email, match).First(&user).Error

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

	return c.JSON(http.StatusOK, map[string]any{
		sortResponse[0]: true,
		sortResponse[1]: "Berhasil Login",
		sortResponse[2]: userResponse,
	})
}

func RegisterAdminController(c echo.Context) error {
	sortResponse := []string{"status", "message", "data"}
	sort.Strings(sortResponse)

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
			sortResponse[0]: false,
			sortResponse[1]: "Email Sudah ada",
			sortResponse[2]: nil,
		})
	}

	//hashing password
	hash, _ := bcrypt.GenerateFromPassword([]byte(userRegister.Password), bcrypt.DefaultCost)

	// userRegister.Password = string(hash)
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
			sortResponse[0]: false,
			sortResponse[1]: err.Error(),
			sortResponse[2]: nil,
		})
    }
	
	if err := config.DB.Model(&user).Create(&newUser).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string] any {
			sortResponse[0]: false,
			sortResponse[1]: "Create failed!",
			sortResponse[2]: nil,
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		sortResponse[0]: true,
		sortResponse[1]: "success create new user",
		sortResponse[2]: newUser,
	})
}

func RegisterUserController(c echo.Context) error {
	sortResponse := []string{"status", "message", "data"}
	// sort.Strings(sortResponse)

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
			sortResponse[0]: false,
			sortResponse[1]: "Email Sudah ada",
			sortResponse[2]: nil,
		})
	}

	//hashing password
	hash, _ := bcrypt.GenerateFromPassword([]byte(userRegister.Password), bcrypt.DefaultCost)

	// userRegister.Password = string(hash)
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
			sortResponse[0]: false,
			sortResponse[1]: err.Error(),
			sortResponse[2]: nil,
		})
    }
	
	if err := config.DB.Model(&user).Create(&newUser).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string] any {
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

func ForgotPasswordController(c echo.Context) error {
	sortResponse := []string{"status", "message", "data"}
	// sort.Strings(sortResponse)
	var users models.User

	var input models.User
	c.Bind(&input)
	email := input.Email

	if err := config.DB.Where("email = ?", email).First(&users).Error; err != nil {
		return c.JSON(http.StatusAlreadyReported, map[string] any {
			sortResponse[0]: false,
			sortResponse[1]: "Email Tidak Ditemukan",
			sortResponse[2]: nil,
		})
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}

	emailTo := email

		data := struct {
			ReceiverName string
			Link 		 string
		}{
			ReceiverName: users.Name,
			Link: "http://github.com/",
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
		sortResponse[0]: true,
		sortResponse[1]: "Sukses, cek emailmu sekarang juga",
		sortResponse[2]: nil,
	})
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}


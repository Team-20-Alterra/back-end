package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"geinterra/config"
	"geinterra/gomail"
	"geinterra/middleware"
	"geinterra/models"
	"geinterra/utils"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/labstack/echo/v4"
	"github.com/thanhpk/randstr"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig *oauth2.Config
	// TODO: randomize it
	oauthStateString = "pseudo-random"
)
func init() {
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  os.Getenv("CALLBACK_URL"),
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

// login
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
			"status":  false,
			"message": "Incorrect Email or Password",
			"data":    nil,
		})
	}

	token, err := middleware.CreateToken(int(user.ID), user.Email, user.Role)
	// token, err := middleware.
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	userResponse := models.UserResponse{int(user.ID), user.Email, user.Role, token}

	return c.JSON(http.StatusOK, map[string]any{
		"status":  true,
		"message": "Berhasil Login",
		"data":    userResponse,
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
			"status":  false,
			"message": "Incorrect Email or Password",
			"data":    nil,
		})
	}

	roleUser := "Admin"

	err = config.DB.Where("role = ?", roleUser).First(&user).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "Only admins can enter",
			"data":    nil,
		})
	}

	token, err := middleware.CreateToken(int(user.ID), user.Email, user.Role)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	userResponse := models.UserResponse{int(user.ID), user.Email, user.Role, token}

	return c.JSON(http.StatusOK, map[string]any{
		"status":  true,
		"message": "Berhasil Login",
		"data":    userResponse,
	})
}

// register
func RegisterAdminController(c echo.Context) error {

	var user models.User
	var userRegister models.UserAdminRegister

	body, _ := ioutil.ReadAll(c.Request().Body)
	err := json.Unmarshal(body, &userRegister)
	if err != nil {
		return err
	}

	email := userRegister.Email

	if err := config.DB.Where("email = ?", email).First(&user).Error; err == nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Email already exist",
			"data":    nil,
		})
	}

	hash, _ := utils.HashPassword(userRegister.Password)

	newUser := models.User{
		Name:          userRegister.Name,
		// Date_of_birth: "",
		Email:         userRegister.Email,
		// Gender:        "",
		// Phone:         userRegister.Phone,
		Address:       "",
		Photo:         "",
		// Username:      "",
		Password:      string(hash),
		Role:          "Admin",
	}

	if err := c.Validate(userRegister); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	if err := config.DB.Model(&user).Create(&newUser).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Create failed!",
			"data":    nil,
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  true,
		"message": "success create new user",
		"data":    newUser,
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
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Email already exist",
			"data":    nil,
		})
	}
	phone := userRegister.Phone

	if err := config.DB.Where("phone = ?", phone).First(&user).Error; err == nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Phone already exist",
			"data":    nil,
		})
	}

	hash, _ := utils.HashPassword(userRegister.Password)

	newUser := models.User{
		Name:          userRegister.Name,
		// Date_of_birth: "",
		Email:         userRegister.Email,
		// Gender:        "",
		Phone:         userRegister.Phone,
		Address:       "",
		Photo:         "",
		// Username:      "",
		Password:      string(hash),
		Role:          "User",
	}

	if err := c.Validate(userRegister); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	if err := config.DB.Model(&user).Create(&newUser).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Create failed!",
			"data":    nil,
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  true,
		"message": "success create new user",
		"data":    newUser,
	})
}
func RegisterBusinessController(c echo.Context) error {
	var busines models.Business
	var business models.BusinessInput
	var user models.User
	var userRegister models.UserAdminRegister
	var list models.LisBankInput

	c.Bind(&list)
	c.Bind(&userRegister)
	c.Bind(&business)
	
	// create user
	email := userRegister.Email

	if err := config.DB.Where("email = ?", email).First(&user).Error; err == nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Email already exist",
			"data":    nil,
		})
	}

	hash, _ := utils.HashPassword(userRegister.Password)

	userRegister.Name = list.Owner
	roleUser := "Admin"
	newUser := models.User{
		Name:          userRegister.Name,
		Email:         userRegister.Email,
		Address:       "",
		Photo:         "",
		Password:      string(hash),
		Role:          roleUser,
	}

	if err := c.Validate(userRegister); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	if err := config.DB.Model(&user).Create(&newUser).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Create failed!",
			"data":    nil,
		})
	}

	// create busies
	// cek already busines
	if err := config.DB.Where("user_id = ?", newUser.ID).First(&busines).Error; err == nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Business already exist",
			"data":    nil,
		})
	}

	// cek user
	if err := config.DB.Where("id = ?", newUser.ID).First(&user).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  false,
			"message": "User not found!",
			"data":    nil,
		})
	}

	// roleUser := "Admin"

	if err := config.DB.Where("role = ?", newUser.Role).First(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "Only admins can create",
			"data":    nil,
		})
	}

	fileHeader, _ := c.FormFile("logo")
	if fileHeader != nil {
		file, _ := fileHeader.Open()

		ctx := context.Background()

		cldService, _ := cloudinary.NewFromURL(os.Getenv("URL_CLOUDINARY"))

		resp, _ := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})

		business.Logo = resp.SecureURL
	}

	fmt.Println(business.No_telp)

	if err := config.DB.Where("no_telp = ?", business.No_telp).First(&busines).Error; err == nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Phone already exist",
			"data":    nil,
		})
	}

	business.UserID = int(newUser.ID)
	busines.Email = newUser.Email

	if err := c.Validate(business); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	businessReal := models.Business{Name: business.Name, Email: busines.Email, Address: business.Address, No_telp: business.No_telp, Type: business.Type, Logo: business.Logo,  UserID: business.UserID}

	if err := config.DB.Create(&businessReal).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}


	// create list bank
	if err := c.Validate(list); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	list.BusinessID = int(businessReal.ID)

	listBank := models.ListBank{Owner: list.Owner, AccountNumber: list.AccountNumber, BankID: list.BankID, BusinessID: list.BusinessID}
	
	if err := config.DB.Create(&listBank).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	token, err := middleware.CreateToken(int(newUser.ID), newUser.Email, newUser.Role)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	var data [4]any

	data  = [4]any{business, list, userRegister}
	
	return c.JSON(http.StatusOK, map[string]any{
		"status":  true,
		"message": "success create new business",
		"data":    data,
		"token": token,
	})
}
// forgot password
func ForgotPasswordController(c echo.Context) error {
	var users models.User

	var input models.ForgotPasswordInput
	c.Bind(&input)
	email := input.Email

	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": err.Error(),
			"data": nil,
		})
	}

	if err := config.DB.Where("email = ?", email).First(&users).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  false,
			"message": "Email Not Found",
			"data": nil,
		})
	}

	// Generate Verification Code
	resetToken := randstr.String(20)

	passwordResetToken := utils.Encode(resetToken)
	
	users.PasswordResetToken = passwordResetToken
	
	users.PasswordResetAt = time.Now().Add(time.Minute * 15)
	
	config.DB.Save(&users)

	emailTo := email

	data := struct{
		ReceiverName string
		Link string
	}{
		ReceiverName: users.Name,
		Link: "https://ginap-mu.vercel.app/new-password?token=" + resetToken,
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
		"status":  true,
		"message": "Success, check your email right now",
		"data":    nil,
	})
}

func ResetPassword(ctx echo.Context) error {
	var payload *models.ResetPasswordInput
	resetToken := ctx.QueryParam("token")

	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": err.Error(),
		})
	}

	if err := ctx.Validate(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	if payload.Password != payload.PasswordConfirm {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Passwords do not match",
		})
	}

	hashedPassword, _ := utils.HashPassword(payload.Password)

	passwordResetToken := utils.Encode(resetToken)

	var updatedUser models.User

	result := config.DB.First(&updatedUser, "password_reset_token = ? AND password_reset_at > ?", 
	
	passwordResetToken, time.Now())
	
	if result.Error != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "The reset token is invalid or has expired",
		})
	}

	updatedUser.Password = hashedPassword
	updatedUser.PasswordResetToken = ""
	
	config.DB.Save(&updatedUser)

	return ctx.JSON(http.StatusOK, map[string]any{
		"status":  true,
		"message": "Password data updated successfully",
	})
}

// login with google
func LoginGoogleController(c echo.Context) error {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(c.Response().Writer, c.Request(), url, http.StatusTemporaryRedirect)

	return c.JSON(200,"ok")
}



func HandleGoogleCallbackController(c echo.Context) error {
	r := c.Request()
	w := c.Response().Writer
	content, err := getUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return c.JSON(500, err.Error())
	}

	var gUser models.GoogleAccount

	err = json.Unmarshal(content, &gUser)

	var user models.User

	// cek ada email / tidak
	if err := config.DB.Where("email = ?", gUser.Email).First(&user).Error; err == nil {
		// jika ada create token
		token, err := middleware.CreateToken(int(user.ID), user.Email, user.Role)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status":  false,
				"message": err.Error(),
				"data":    nil,
			})
		}		
		
		userResponse := models.UserResponse{int(user.ID), user.Email, user.Role, token}

		return c.JSON(http.StatusOK, map[string]any{
			"status": true,
			"message": "success login google",
			"data": userResponse,
		})
	}

	// kalau tidak create user & create token
	hash, _ := utils.HashPassword("Password")

	newUser := models.User{
		Name:          gUser.Email,
		Email:         gUser.Email,
		Phone:         "",
		Address:       "",
		Photo:         "",
		Password:      string(hash),
		Role:          "User",
	}

	if err := config.DB.Model(&user).Create(&newUser).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Create failed!",
			"data":    nil,
		})
	}

	token, err := middleware.CreateToken(int(newUser.ID), newUser.Email, newUser.Role)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  false,
			"message": err.Error(),
			"data":    nil,
		})
	}		
	
	userResponse := models.UserResponse{int(newUser.ID), newUser.Email, newUser.Role, token}
	return c.JSON(200,map[string]any{
		"status": true,
		"message": "success auth google",
		"data": userResponse,
	})
}

func getUserInfo(state string, code string) ([]byte, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	fmt.Println(contents)

	return contents, nil
}
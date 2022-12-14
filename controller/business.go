package controller

import (
	"context"
	"geinterra/config"
	"geinterra/models"
	"net/http"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetBusinesssController(c echo.Context) error {
	var business []models.BusinessResponse

	if err := config.DB.Model(&models.Business{}).Joins("User").Select("businesses.id,businesses.name,businesses.email,businesses.address,businesses.no_telp,businesses.type,businesses.logo,businesses.user_id,User.phone,User.address,User.photo").Scan(&business).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Busines not found!",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "success get all business",
		"data":    business,
	})
}

func GetBusinessController(c echo.Context) error {
	var busines models.Business
	var business models.BusinessResponse

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Model(&models.Business{}).Joins("User").Select("businesses.id,businesses.name,businesses.email,businesses.address,businesses.no_telp,businesses.type,businesses.logo,businesses.user_id,User.phone,User.address,User.photo").Where("businesses.id = ?", id).First(&busines).Scan(&business).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Busines not found!",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "success get business",
		"data":    business,
	})
}

func GetBusinessByUserController(c echo.Context) error {
	var busines models.Business
	var business models.BusinessResponse

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

	if err := config.DB.Model(&models.Business{}).Joins("User").Select("businesses.id,businesses.name,businesses.email,businesses.address,businesses.no_telp,businesses.type,businesses.logo,businesses.user_id,User.phone,User.address,User.photo").Where("businesses.user_id = ?", id).First(&busines).Scan(&business).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Busines not found!",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "success get business",
		"data":    business,
	})
}

func UpdateBusinessController(c echo.Context) error {
	var busines models.Business

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

	// cek already busines
	if err := config.DB.Where("user_id = ?", id).First(&busines).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Business already exist",
			"data":    nil,
		})
	}

	var input models.BusinessUpdate
	c.Bind(&input)

	fileHeader, _ := c.FormFile("logo")
	if fileHeader != nil {
		file, _ := fileHeader.Open()

		ctx := context.Background()

		cldService, _ := cloudinary.NewFromURL("cloudinary://852912385417941:-GFfGWwjDwrsPgyH7ZMXEvuc9DM@dwdaw6znj")

		resp, _ := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})

		input.Logo = resp.SecureURL
	}

	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	// validate busines
	if err := config.DB.Where("id = ?", busines.ID).First(&busines).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  false,
			"message": "Busines not found!",
			"data":    nil,
		})
	}

	businessReal := models.Business{Name: input.Name, Address: input.Address, No_telp: input.No_telp, Type: input.Type, Email: input.Email, Logo: input.Logo}

	if err := config.DB.Model(&busines).Where("id = ?", busines.ID).Updates(businessReal).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  false,
			"message": "Record not found!",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "update success",
	})
}

func DeleteBusinessController(c echo.Context) error {
	var busines models.Business

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

	// cek already busines
	if err := config.DB.Where("user_id = ?", id).First(&busines).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Business already exist",
			"data":    nil,
		})
	}

	// validate busines
	if err := config.DB.Where("id = ?", busines.ID).First(&busines).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  false,
			"message": "Busines not found!",
			"data":    nil,
		})
	}

	if err := config.DB.Delete(&busines, busines.ID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  false,
			"message": "Record not found!",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "success delete Business",
	})
}

// func CreateBusinessController(c echo.Context) error {
// 	var users models.User
// 	var busines models.Business
// 	var business models.BusinessInput

// 	c.Bind(&business)

// 	user := c.Get("user").(*jwt.Token)
// 	claims := user.Claims.(jwt.MapClaims)

// 	id, _ := claims["id"]

// 	userId := id.(float64)

// 	// cek already busines
// 	if err := config.DB.Where("user_id = ?", id).First(&busines).Error; err == nil {
// 		return c.JSON(http.StatusBadRequest, map[string]any{
// 			"status":  false,
// 			"message": "Business already exist",
// 			"data":    nil,
// 		})
// 	}

// 	// cek user
// 	if err := config.DB.Where("id = ?", id).First(&users).Error; err != nil {
// 		return c.JSON(http.StatusNotFound, map[string]any{
// 			"status":  false,
// 			"message": "User not found!",
// 			"data":    nil,
// 		})
// 	}

// 	roleUser := "Admin"

// 	if err := config.DB.Where("role = ?", roleUser).First(&user).Error; err == nil {
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"status":  false,
// 			"message": "Only admins can create",
// 			"data":    nil,
// 		})
// 	}

// 	fileHeader, _ := c.FormFile("logo")
// 	if fileHeader != nil {
// 		file, _ := fileHeader.Open()

// 		ctx := context.Background()

// 		cldService, _ := cloudinary.NewFromURL(os.Getenv("URL_CLOUDINARY"))

// 		resp, _ := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})

// 		business.Logo = resp.SecureURL
// 	}

// 	business.UserID = int(userId)

// 	if err := c.Validate(business); err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	businessReal := models.Business{Name: business.Name, Address: business.Address, No_telp: business.No_telp, Type: business.Type, Logo: business.Logo,  UserID: business.UserID}

// 	if err := config.DB.Create(&businessReal).Error; err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err)
// 	}

// 	// create listbank
// 	var list models.LisBankInput

// 	c.Bind(&list)

// 	if err := c.Validate(list); err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	list.BusinnesID = int(businessReal.ID)

// 	listBank := models.ListBank{Owner: list.Owner, AccountNumber: list.AccountNumber, BankID: list.BankID, BusinnesID: list.BusinnesID}

// 	if err := config.DB.Create(&listBank).Error; err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err)
// 	}

// 	var data [2]any

// 	data  = [2]any{business, list}

// 	return c.JSON(http.StatusOK, map[string]any{
// 		"status":  true,
// 		"message": "success create new business",
// 		"data":    data,
// 	})
// }

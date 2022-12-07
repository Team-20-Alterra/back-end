package controller

import (
	"context"
	"encoding/json"
	"geinterra/config"
	"geinterra/models"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetBusinesssController(c echo.Context) error {
	var business []models.Business

	if err := config.DB.Preload("Bank").Find(&business).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success get all business",
		"Business": business,
	})
}

func GetBusinessController(c echo.Context) error {
	var business models.Business

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Where("id = ?", id).Preload("Bank").First(&business).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success get business",
		"business": business,
	})
}

func CreateBusinessController(c echo.Context) error {
	var users models.User
	var busines models.Business
	var business models.BusinessInput
<<<<<<< HEAD

	body, _ := ioutil.ReadAll(c.Request().Body)
	err := json.Unmarshal(body, &business)
	if err != nil {
		return err
	}
=======
	c.Bind(&business)
>>>>>>> d1ff30ce34593877e7beda68e4572bac178c8241

	fileHeader, _ := c.FormFile("logo")
	if fileHeader != nil {
		file, _ := fileHeader.Open()

		ctx := context.Background()

		cldService, _ := cloudinary.NewFromURL(os.Getenv("URL_CLOUDINARY"))

		resp, _ := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})

		business.Logo = resp.SecureURL
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

	userId := id.(float64)

	// cek already busines
	if err := config.DB.Where("user_id = ?", id).First(&busines).Error; err == nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  false,
			"message": "Business already exist",
			"data":    nil,
		})
	}

	// cek user
	if err := config.DB.Where("id = ?", id).First(&users).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  false,
			"message": "User not found!",
			"data":    nil,
		})
	}

	roleUser := "Admin"

	if err := config.DB.Where("role = ?", roleUser).First(&user).Error; err == nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "Only admins can create",
			"data":    nil,
		})
	}

	if err := c.Validate(business); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	businessReal := models.Business{Name: business.Name, Address: business.Address, No_telp: business.No_telp, Type: business.Type, Logo: business.Logo, BankID: business.BankID, UserID: int(userId)}

	if err := config.DB.Create(&businessReal).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// create listbank
	var list models.LisBankInput

	err2 := json.Unmarshal(body, &list)
	if err2 != nil {
		return err
	}

	if err := c.Validate(list); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	list.BusinnesID = int(businessReal.ID)

	listBank := models.ListBank{Owner: list.Owner, AccountNumber: list.AccountNumber, BankID: list.BankID, BusinnesID: list.BusinnesID}
	
	if err := config.DB.Create(&listBank).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	var data [2]any

<<<<<<< HEAD
	data  = [2]any{business, list}
	
=======
	// data  = [2]string{business, busines}

>>>>>>> d1ff30ce34593877e7beda68e4572bac178c8241
	return c.JSON(http.StatusOK, map[string]any{
		"status":  true,
		"message": "success create new business",
		"data":    data,
	})
}

func UpdateBusinessController(c echo.Context) error {
	sortResponse := []string{"status", "message", "data"}
	sort.Strings(sortResponse)

	id, _ := strconv.Atoi(c.Param("id"))

	var business models.Business

	var input models.Business
	c.Bind(&input)

	fileHeader, _ := c.FormFile("logo")
	if fileHeader != nil {
		file, _ := fileHeader.Open()

		ctx := context.Background()

		cldService, _ := cloudinary.NewFromURL(os.Getenv("URL_CLOUDINARY"))

		resp, _ := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})

		input.Logo = resp.SecureURL
	}

	if err := config.DB.Model(&business).Where("id = ?", id).Updates(input).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			sortResponse[0]: false,
			sortResponse[1]: "Record not found!",
			sortResponse[2]: nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		sortResponse[0]: true,
		sortResponse[1]: "update success",
	})
}

func DeleteBusinessController(c echo.Context) error {
	sortResponse := []string{"status", "message", "data"}
	sort.Strings(sortResponse)

	var business models.Business

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Delete(&business, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			sortResponse[0]: false,
			sortResponse[1]: "Record not found!",
			sortResponse[2]: nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		sortResponse[0]: true,
		sortResponse[1]: "success delete Business",
	})
}

package controller

import (
	"context"
	"geinterra/config"
	"geinterra/models"
	"net/http"
	"os"
	"sort"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/labstack/echo/v4"
)

func GetBanksController(c echo.Context) error {
	var bank []models.Bank

	if err := config.DB.Find(&bank).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success get all bank",
		"bank":   bank,
	})
}

func GetBankController(c echo.Context) error {
	var bank models.Bank

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Where("id = ?", id).First(&bank).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get bank",
		"bank":    bank,
	})
}

func CreateBankController(c echo.Context) error {
	var bank models.Bank
	c.Bind(&bank)

	fileHeader, _ := c.FormFile("logo")
	if fileHeader != nil {
		file, _ := fileHeader.Open()

		ctx := context.Background()

		cldService, _ := cloudinary.NewFromURL(os.Getenv("URL_CLOUDINARY"))

		resp, _ := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})

		bank.Logo = resp.SecureURL
	}

	if err := c.Validate(bank); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Create(&bank).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "success create new bank",
		"data":    bank,
	})
}

func UpdateBankController(c echo.Context) error {
	sortResponse := []string{"status", "message", "data"}
	sort.Strings(sortResponse)

	id, _ := strconv.Atoi(c.Param("id"))

	var bank models.Bank

	var input models.Bank
	c.Bind(&input)

	fileHeader, _ := c.FormFile("logo")
	if fileHeader != nil {
		file, _ := fileHeader.Open()

		ctx := context.Background()

		cldService, _ := cloudinary.NewFromURL(os.Getenv("URL_CLOUDINARY"))

		resp, _ := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})

		input.Logo = resp.SecureURL
	}

	if err := config.DB.Model(&bank).Where("id = ?", id).Updates(input).Error; err != nil {
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

func DeleteBankController(c echo.Context) error {
	sortResponse := []string{"status", "message", "data"}
	sort.Strings(sortResponse)

	var bank models.Bank

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Delete(&bank, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			sortResponse[0]: false,
			sortResponse[1]: "Record not found!",
			sortResponse[2]: nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		sortResponse[0]: true,
		sortResponse[1]: "success delete bank",
	})
}

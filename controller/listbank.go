package controller

import (
	"geinterra/config"
	"geinterra/models"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetListBanksController(c echo.Context) error {
	var list []models.LisBankResponse

	if err := config.DB.Model(&models.ListBank{}).Joins("Bank").Select("list_banks.id,list_banks.owner,list_banks.account_number,list_banks.bank_id,list_banks.business_id").Scan(&list).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  false,
			"message": "List Bank not found!",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "success get all listBank",
		"data":    list,
	})
}

func GetListBankByIdController(c echo.Context) error {
	var list models.LisBankResponse
	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Model(&models.ListBank{}).Joins("Bank").Select("list_banks.id,list_banks.owner,list_banks.account_number,list_banks.bank_id,list_banks.business_id").Where("list_banks.id = ?", id).Scan(&list).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  false,
			"message": "List Bank not found!",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "success get listBank",
		"data":    list,
	})
}

func GetListBankByBusinessController(c echo.Context) error {
	// var listBank []models.ListBank
	var list []models.LisBankResponse
	var busines models.Business

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

	if err := config.DB.Joins("JOIN list_banks on list_banks.business_id=businesses.id").
		Where("businesses.user_id=?", id).
		Group("businesses.id").Preload("User").Find(&busines).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	if busines.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  false,
			"message": "Record not found!",
			"data":    nil,
		})
	}

	if err := config.DB.Model(&models.ListBank{}).Joins("Bank").Select("list_banks.id,list_banks.owner,list_banks.account_number,list_banks.bank_id,list_banks.business_id").Where("business_id = ?", busines.ID).Scan(&list).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  false,
			"message": "List Bank not found!",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "success get listBank",
		"data":    list,
	})
}

func CreateListBankController(c echo.Context) error {
	var list models.LisBankInput
	var lisBank models.ListBank
	var busines models.Business
	var bank models.Bank

	c.Bind(&list)

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["id"]

	if err := config.DB.Joins("JOIN list_banks on list_banks.business_id=businesses.id").
		Where("businesses.user_id=?", id).
		Group("businesses.id").Preload("User").Find(&busines).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	if busines.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  false,
			"message": "Record not found!",
			"data":    nil,
		})
	}

	// cek bank
	if err := config.DB.Where("id = ?", list.BankID).First(&bank).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  false,
			"message": "Bank Not Found",
			"data":    nil,
		})
	}

	// cek already list bank
	if err := config.DB.Where("bank_id = ? AND business_id = ?", list.BankID, busines.ID).First(&lisBank).Error; err == nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  false,
			"message": "list bank already",
			"data":    nil,
		})
	}

	if err := c.Validate(list); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	list.BusinessID = int(busines.ID)

	listBank := models.ListBank{Owner: list.Owner, AccountNumber: list.AccountNumber, BankID: list.BankID, BusinessID: list.BusinessID}

	if err := config.DB.Create(&listBank).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status":  true,
		"message": "success create new business",
		"data":    list,
	})
}

package middleware

import (
	"geinterra/constants"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateToken(userId int, username string, email string, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["username"] = username
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_KEY))
}

var whitelist []string = make([]string, 5)

type JwtCustomClaims struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Role string `json:"role"`
	jwt.StandardClaims
}

func GetUser(c echo.Context) *JwtCustomClaims {
	user := c.Get("user").(*jwt.Token)

	isListed := CheckToken(user.Raw)

	if !isListed {
		return nil
	}

	claims := user.Claims.(*JwtCustomClaims)
	return claims
}

func CheckToken(token string) bool {
	for _, tkn := range whitelist {
		if tkn == token {
			return true
		}
	}

	return false
}
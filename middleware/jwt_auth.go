package middleware

import (
	"geinterra/constants"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(userId int, username string, status string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["username"] = username
	claims["status"] = status
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_KEY))
}

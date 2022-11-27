package middleware

import (
	"geinterra/constants"
	"time"

	"github.com/golang-jwt/jwt"
)

var whitelist []string = make([]string, 5)

type JwtCustomClaims struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Role string `json:"role"`
	jwt.StandardClaims
}
type JwtCustomClaimsForgot struct {
	ID int `json:"id"`
	Password string `json:"password"`
	jwt.StandardClaims
}


func CreateToken(userId int, username string, email string, role string) (string, error) {
	claims := &JwtCustomClaims{
		userId,
		username,
		email,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_KEY))
}

func CreateTokenForgot(userId int, password string) (string, error) {
	claims := &JwtCustomClaimsForgot{
		userId,
		password,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_KEY))
}
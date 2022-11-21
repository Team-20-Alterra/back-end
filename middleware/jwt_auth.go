package middleware

import (
	"geinterra/constants"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var whitelist []string = make([]string, 5)

type JwtCustomClaims struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Role string `json:"role"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT       string
	ExpiresDuration int
}

func CreateToken(userId int, username string, email string, role string) (string, error) {
	// claims := jwt.MapClaims{}
	// claims["userId"] = userId
	// claims["username"] = username
	// claims["role"] = role
	// claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	// Set custom claims
	claims := &JwtCustomClaims{
		userId,
		username,
		email,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	// t, err := token.SignedString([]byte("secret"))

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_KEY))
}



func (jwtConf *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
	}
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
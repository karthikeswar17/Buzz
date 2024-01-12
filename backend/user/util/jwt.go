package util

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var SigningKey = "Secret"

func CreateJWT(userId string, email string, tag string) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["user-id"] = userId
	claims["tag"] = tag
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response
	encToken, err := token.SignedString([]byte(SigningKey))
	if err != nil {
		return "", err
	}
	return encToken, nil

}

func GetEmailAndTagFromJWT(c echo.Context) (string, string) {
	user, _ := c.Get("user").(*jwt.Token)
	claims, _ := user.Claims.(jwt.MapClaims)
	return claims["email"].(string), claims["tag"].(string)
}

func GetUserIdFromJWT(c echo.Context) string {
	user, _ := c.Get("user").(*jwt.Token)
	claims, _ := user.Claims.(jwt.MapClaims)
	return claims["user-id"].(string)
}

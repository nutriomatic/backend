package middleware

import (
	"fmt"
	"golang-template/models"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func GenerateTokenPair(user *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = 1
	claims["name"] = user.Name
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return t, nil
}

func GetToken(c echo.Context) string {
	auth, exists := c.Request().Header["Authorization"]
	if !exists || len(auth) == 0 {
		fmt.Println("Error: Authorization header is missing or empty")
		return ""
	}

	Bearer := auth[0]
	tokenParts := strings.Split(Bearer, "Bearer ")
	if len(tokenParts) != 2 {
		fmt.Println("Error: Invalid Bearer token format")
		return ""
	}

	token := tokenParts[1]
	return token
}

func GetTokenNext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		if auth == "" {
			return c.String(http.StatusUnauthorized, "Authorization header is missing")
		}

		parts := strings.Split(auth, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.String(http.StatusUnauthorized, "Invalid token format")
		}

		token := parts[1]
		c.Set("token", token) // Store token in the context for further use
		return next(c)
	}
}

package middleware

import (
	"fmt"
	"golang-template/models"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
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
		return "Error in jwt", err
	}

	return t, nil
}

func GetToken(c echo.Context) string {
	auth, exists := c.Request().Header["Authorization"]
	if !exists || len(auth) == 0 {
		fmt.Println("Error: Authorization header is missing or empty")
		return "Error in jwt"
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
		// Verify the JWT token
		claims := jwt.MapClaims{}
		t, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			// Check the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("secret"), nil // Specify your secret key here
		})
		if err != nil {
			return c.String(http.StatusUnauthorized, "Invalid token")
		}
		if !t.Valid {
			return c.String(http.StatusUnauthorized, "Invalid token")
		}

		// Check the audience claim (aud)
		audience, ok := claims["aud"].([]string)
		if !ok || len(audience) == 0 {
			return c.String(http.StatusUnauthorized, "Invalid audience claim")
		}

		// Store token in the context for further use
		c.Set("token", token)
		return next(c)
	}
}

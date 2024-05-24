package controllers

import (
	"golang-template/dto"
	"golang-template/middleware"
	"golang-template/models"
	"golang-template/services"
	"golang-template/utils"
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type authController struct {
	userService  services.UserService
	tokenService services.TokenService
}

func NewAuthController() *authController {
	return &authController{
		userService:  services.NewUserService(),
		tokenService: services.NewTokenService(),
	}
}

var log = logrus.New()

func init() {
	log.SetLevel(logrus.DebugLevel)
}

func (a *authController) Register(c echo.Context) error {
	registerReq := &dto.RegisterForm{}
	err := c.Bind(registerReq)
	if err != nil {
		return c.String(http.StatusBadRequest, "All user fields must be provided!")
	}

	_, err = a.userService.GetUserByUsername(registerReq.Username)
	if err == nil {
		return c.String(http.StatusBadRequest, "Username "+registerReq.Username+" already exists.")
	}

	_, err = a.userService.GetUserByEmail(registerReq.Email)
	if err == nil {
		return c.String(http.StatusBadRequest, "Email "+registerReq.Email+" already exists")
	}

	if !utils.ValidateLengthPassword(registerReq.Password) {
		return c.String(http.StatusBadRequest, "Password is too short")
	}

	err = a.userService.CreateUser(registerReq)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"status":  "success",
		"message": dto.Register_Successful,
	})
}

func (a *authController) Login(c echo.Context) error {
	loginReq := &dto.LoginForm{}
	err := c.Bind(loginReq)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid login form.")
	}

	var user *models.User
	if loginReq.Username == "" {
		user, err = a.userService.GetUserByEmail(loginReq.Email)
	} else {
		user, err = a.userService.GetUserByUsername(loginReq.Username)
	}
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error retrieving user.")
	}

	if user == nil {
		return c.String(http.StatusUnauthorized, "User not found.")
	}

	if utils.ValidatePassword(user.Password, loginReq.Password) {
		token, err := middleware.GenerateTokenPair(user)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error generating token.")
		}
		err = a.tokenService.SaveToken(user, token)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error saving token.")
		}
		var loginmessage string

		if user.Role == "customer" {
			loginmessage = "You are logged in as customer"
		} else {
			loginmessage = "You are logged in as admin"
		}

		return c.JSON(http.StatusOK, map[string]string{
			"status":  "success",
			"message": loginmessage,
			"token":   token,
		})
	}

	return echo.ErrUnauthorized
}

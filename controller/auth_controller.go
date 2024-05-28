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
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "failed",
			"message": dto.FieldsRequired,
		})
	}

	_, err = a.userService.GetUserByUsername(registerReq.Username)
	if err == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "failed",
			"message": dto.UsernameExists,
		})
	}

	_, err = a.userService.GetUserByEmail(registerReq.Email)
	if err == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "failed",
			"message": dto.EmailExists,
		})
	}

	if !utils.ValidateLengthPassword(registerReq.Password) {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "failed",
			"message": dto.PasswordShort,
		})
	}

	err = a.userService.CreateUser(registerReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "failed",
			"message": "Error creating user.",
		})
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
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "failed",
			"message": dto.ErrorRetrievingUser,
		})
	}

	if user == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"status":  "failed",
			"message": dto.UserNotFound,
		})
	}

	if utils.ValidatePassword(user.Password, loginReq.Password) {
		token, err := middleware.GenerateTokenPair(user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"status":  "failed",
				"message": dto.ErrorGeneratingToken,
			})
		}
		err = a.tokenService.SaveToken(user, token)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"status":  "failed",
				"message": dto.ErrorSavingToken,
			})
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

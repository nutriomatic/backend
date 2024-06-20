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
	registerReq := &dto.Register{}
	err := c.Bind(registerReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":    "400",
			"status":  "failed",
			"message": dto.FieldsRequired,
		})
	}

	_, err = a.userService.GetUserByEmail(registerReq.Email)
	if err == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":    "400",
			"status":  "failed",
			"message": dto.EmailExists,
		})
	}

	if !utils.ValidateLengthPassword(registerReq.Password) {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":    "400",
			"status":  "failed",
			"message": dto.PasswordShort,
		})
	}

	err = a.userService.CreateUser(registerReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"status":  "failed",
			"message": "Error creating user.",
		})
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"code":    "201",
		"status":  "success",
		"message": dto.Register_Successful,
	})
}

func (a *authController) Login(c echo.Context) error {
	loginReq := &dto.LoginForm{}
	err := c.Bind(loginReq)
	if err != nil {
		response := map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": dto.FieldsRequired,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	var user *models.User
	user, err = a.userService.GetUserByEmail(loginReq.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"status":  "failed",
			"message": dto.ErrorRetrievingUser,
		})
	}

	if user == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"code":    "401",
			"status":  "failed",
			"message": dto.UserNotFound,
		})
	}

	if utils.ValidatePassword(user.Password, loginReq.Password) {
		token, err := middleware.GenerateTokenPair(user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"code":    "500",
				"status":  "failed",
				"message": dto.ErrorGeneratingToken,
			})
		}
		err = a.tokenService.SaveToken(user, token)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"code":    "500",
				"status":  "failed",
				"message": dto.ErrorSavingToken,
			})
		}
		var loginmessage string

		if user.Role == "admin" {
			loginmessage = "You are logged in as admin"
		} else {
			loginmessage = "You are logged in as customer"
		}

		return c.JSON(http.StatusOK, map[string]string{
			"code":    "200",
			"status":  "success",
			"role":    user.Role,
			"message": loginmessage,
			"token":   token,
		})
	}

	return echo.ErrUnauthorized
}

func (a *authController) CheckToken(c echo.Context) error {
	token := middleware.GetToken(c)
	return a.tokenService.CheckToken(c, token)
}

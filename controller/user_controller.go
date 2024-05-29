package controllers

import (
	"golang-template/dto"
	"golang-template/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userController struct {
	UserService  services.UserService
	TokenService services.TokenService
}

func NewUserController() *userController {
	return &userController{
		UserService:  services.NewUserService(),
		TokenService: services.NewTokenService(),
	}
}

func (u *userController) GetUserByToken(c echo.Context) error {
	user, err := u.TokenService.UserByToken(c)
	if err != nil {
		return echo.ErrUnauthorized
	} else {
		response := map[string]interface{}{
			"status": "success",
			"user":   user,
		}
		return c.JSON(http.StatusOK, response)
	}
}

func (u *userController) UpdateUser(c echo.Context) error {
	updateForm := &dto.RegisterForm{}
	err := c.Bind(updateForm)
	if err != nil {
		response := map[string]interface{}{
			"status":  "error",
			"message": "All user fields must be provided!",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	updatedUser, err := u.UserService.UpdateUser(updateForm, c)
	if err != nil {
		httpError, ok := err.(*echo.HTTPError)
		if ok {
			response := map[string]interface{}{
				"status":  "error",
				"message": httpError.Message,
			}
			return c.JSON(httpError.Code, response)
		}
		response := map[string]interface{}{
			"status":  "error",
			"message": "internal server error",
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "user updated",
		"user":    updatedUser,
	}
	return c.JSON(http.StatusOK, response)
}

func (u *userController) DeleteUser(c echo.Context) error {
	err := u.UserService.DeleteUser(c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "error in removing user")
	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "account removed successfully",
	}
	return c.JSON(http.StatusOK, response)
}

func (u *userController) Logout(c echo.Context) error {
	err := u.UserService.Logout(c)
	if err != nil {
		return err
	} else {
		response := map[string]interface{}{
			"status":  "success",
			"message": "Logout was successful",
		}
		return c.JSON(http.StatusOK, response)
	}
}

package controllers

import (
	"golang-template/dto"
	"golang-template/middleware"
	"golang-template/repository"
	"golang-template/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userController struct {
	UserService  services.UserService
	TokenService services.TokenService
	TokenRepo    repository.TokenRepository
}

func NewUserController() *userController {
	return &userController{
		UserService:  services.NewUserService(),
		TokenService: services.NewTokenService(),
		TokenRepo:    repository.NewTokenRepositoryGORM(),
	}
}

func (u *userController) GetUserById(c echo.Context) error {
	tokenUser, err := u.TokenRepo.UserToken(middleware.GetToken(c))
	if err != nil {
		response := map[string]interface{}{
			"code":    http.StatusUnauthorized,
			"status":  "error",
			"message": "Unauthorized",
		}
		return c.JSON(http.StatusUnauthorized, response)
	}

	user, err := u.UserService.GetUserById(tokenUser.ID)
	if err != nil {
		response := map[string]interface{}{
			"code":    http.StatusNotFound,
			"status":  "error",
			"message": "User not found",
		}
		return c.JSON(http.StatusNotFound, response)
	}
	response := map[string]interface{}{
		"code":   http.StatusOK,
		"status": "success",
		"user":   user,
	}
	return c.JSON(http.StatusOK, response)
}

func (u *userController) GetUserByToken(c echo.Context) error {
	user, err := u.TokenService.UserByToken(c)
	if err != nil {
		return echo.ErrUnauthorized
	} else {
		response := map[string]interface{}{
			"code":   http.StatusOK,
			"status": "success",
			"user":   user,
		}
		return c.JSON(http.StatusOK, response)
	}
}

func (u *userController) GetClassCalories(c echo.Context) error {
	calories, classification, err := u.UserService.GetClassCalories(c)
	if err != nil {
		response := map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"status":  "error",
			"message": "internal server error",
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := map[string]interface{}{
		"code":           http.StatusOK,
		"status":         "success",
		"calories":       calories,
		"classification": classification,
	}
	return c.JSON(http.StatusOK, response)
}

func (u *userController) UpdateUser(c echo.Context) error {
	updateForm := &dto.RegisterForm{}
	err := c.Bind(updateForm)
	if err != nil {
		response := map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": "All user fields must be provided!",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	err = u.UserService.UpdateUser(updateForm, c)
	if err != nil {
		httpError, ok := err.(*echo.HTTPError)
		if ok {
			response := map[string]interface{}{
				"code":    httpError.Code,
				"status":  "error",
				"message": httpError.Message,
			}
			return c.JSON(httpError.Code, response)
		}
		response := map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"status":  "error",
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := map[string]interface{}{
		"code":    http.StatusOK,
		"status":  "success",
		"message": "user updated",
	}
	return c.JSON(http.StatusOK, response)
}

func (u *userController) DeleteUser(c echo.Context) error {
	err := u.UserService.DeleteUser(c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "error in removing user")
	}

	response := map[string]interface{}{
		"code":    http.StatusOK,
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
			"code":    http.StatusOK,
			"status":  "success",
			"message": "Logout was successful",
		}
		return c.JSON(http.StatusOK, response)
	}
}

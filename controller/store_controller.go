package controllers

import (
	"golang-template/dto"
	"golang-template/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type storeController struct {
	StoreService services.StoreService
	TokenService services.TokenService
}

func NewStoreController() *storeController {
	return &storeController{
		StoreService: services.NewStoreService(),
		TokenService: services.NewTokenService(),
	}
}

func (s *storeController) CreateStore(c echo.Context) error {
	user, err := s.TokenService.UserToken(c)
	if err != nil {
		return echo.ErrUnauthorized
	} else {
		StoreForm := &dto.StoreRegisterForm{}
		err := c.Bind(StoreForm)
		if err != nil {
			response := map[string]interface{}{
				"status":  "error",
				"message": "All store fields must be provided!",
			}
			return c.JSON(http.StatusBadRequest, response)
		}

		err = s.StoreService.CreateStore(StoreForm, user)
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
			"message": "store registration was successful",
		}
		return c.JSON(http.StatusOK, response)
	}
}

func (s *storeController) GetStoreByUserId(c echo.Context) error {
	user, err := s.TokenService.UserToken(c)
	if err != nil {
		return echo.ErrUnauthorized
	} else {
		store, err := s.StoreService.GetStoreByUserId(user.ID)
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
		return c.JSON(http.StatusOK, store)
	}
}

func (s *storeController) UpdateStore(c echo.Context) error {
	updateForm := &dto.StoreRegisterForm{}
	err := c.Bind(updateForm)
	if err != nil {
		response := map[string]interface{}{
			"status":  "error",
			"message": "All user fields must be provided!",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	err = s.StoreService.UpdateStore(updateForm, c)
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
		"message": "successfully updated store",
	}
	return c.JSON(http.StatusOK, response)
}

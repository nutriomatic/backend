package controllers

import (
	"golang-template/dto"
	"golang-template/services"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type HealthGoalController struct {
	HealthGoalService services.HealthGoalService
	TokenService      services.TokenService
}

func NewHealthGoalController() *HealthGoalController {
	return &HealthGoalController{
		HealthGoalService: services.NewHealthGoalService(),
		TokenService:      services.NewTokenService(),
	}
}

func (hg *HealthGoalController) CreateHealthGoal(c echo.Context) error {
	if !hg.TokenService.IsAdmin(c) {
		response := map[string]interface{}{
			"status":  "failed",
			"message": "Unauthorized access",
		}

		return c.JSON(http.StatusUnauthorized, response)
	}

	HGForm := &dto.HGRegisterForm{}
	err := c.Bind(HGForm)
	if err != nil {
		response := map[string]interface{}{
			"status":  "error",
			"message": "All health goal fields must be provided!",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	err = hg.HealthGoalService.CreateHealthGoal(HGForm)
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
		"message": "health goal registration was successful",
	}
	return c.JSON(http.StatusOK, response)
}

func (hg *HealthGoalController) DeleteHealthGoal(c echo.Context) error {
	if !hg.TokenService.IsAdmin(c) {
		response := map[string]interface{}{
			"status":  "failed",
			"message": "Unauthorized access",
		}

		return c.JSON(http.StatusUnauthorized, response)
	}

	id := c.Param("id")
	err := hg.HealthGoalService.DeleteHealthGoal(id)
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
		"message": "health goal deletion was successful",
	}
	return c.JSON(http.StatusOK, response)
}

func (hg *HealthGoalController) GetAllHealthGoal(c echo.Context) error {
	if !hg.TokenService.IsAdmin(c) {
		response := map[string]interface{}{
			"status":  "failed",
			"message": "Unauthorized access",
		}

		return c.JSON(http.StatusUnauthorized, response)
	}

	hgs, err := hg.HealthGoalService.GetAllHealthGoal()
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
		"status": "success",
		"data":   hgs,
	}
	return c.JSON(http.StatusOK, response)
}

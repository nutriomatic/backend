package controllers

import (
	"golang-template/dto"
	"golang-template/services"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type ActivityLevelController struct {
	ActivityLevelService services.ActivityLevelService
	TokenService         services.TokenService
}

func NewActivityLevelController() *ActivityLevelController {
	return &ActivityLevelController{
		ActivityLevelService: services.NewActivityLevelService(),
		TokenService:         services.NewTokenService(),
	}
}

func (al *ActivityLevelController) CreateActivityLevel(c echo.Context) error {
	if !al.TokenService.IsAdmin(c) {
		response := map[string]interface{}{
			"status":  "failed",
			"message": "Unauthorized access",
		}

		return c.JSON(http.StatusUnauthorized, response)
	}

	ALForm := &dto.ALRegisterForm{}
	err := c.Bind(ALForm)
	if err != nil {
		response := map[string]interface{}{
			"status":  "error",
			"message": "All activity level fields must be provided!",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	err = al.ActivityLevelService.CreateAL(ALForm)
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
		"message": "Activity level created successfully",
	}
	return c.JSON(http.StatusCreated, response)
}

func (al *ActivityLevelController) GetAllActivityLevel(c echo.Context) error {
	if !al.TokenService.IsAdmin(c) {
		response := map[string]interface{}{
			"status":  "failed",
			"message": "Unauthorized access",
		}

		return c.JSON(http.StatusUnauthorized, response)
	}

	activityLevels, err := al.ActivityLevelService.GetAllActivityLevel()
	if err != nil {
		response := map[string]interface{}{
			"status":  "error",
			"message": "internal server error",
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Activity levels retrieved successfully",
		"data":    activityLevels,
	}
	return c.JSON(http.StatusOK, response)
}

func (al *ActivityLevelController) GetActivityLevelById(c echo.Context) error {
	if !al.TokenService.IsAdmin(c) {
		response := map[string]interface{}{
			"status":  "failed",
			"message": "Unauthorized access",
		}

		return c.JSON(http.StatusUnauthorized, response)
	}

	id := c.Param("id")
	activityLevel, err := al.ActivityLevelService.GetById(id)
	if err != nil {
		response := map[string]interface{}{
			"status":  "error",
			"message": "Activity level not found",
		}
		return c.JSON(http.StatusNotFound, response)
	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Activity level retrieved successfully",
		"data":    activityLevel,
	}
	return c.JSON(http.StatusOK, response)
}

func (al *ActivityLevelController) DeleteActivityLevel(c echo.Context) error {
	if !al.TokenService.IsAdmin(c) {
		response := map[string]interface{}{
			"status":  "failed",
			"message": "Unauthorized access",
		}

		return c.JSON(http.StatusUnauthorized, response)
	}

	id := c.Param("id")
	err := al.ActivityLevelService.DeleteAL(id)
	if err != nil {
		response := map[string]interface{}{
			"status":  "error",
			"message": "Activity level not found",
		}
		return c.JSON(http.StatusNotFound, response)
	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Activity level deleted successfully",
	}
	return c.JSON(http.StatusOK, response)
}

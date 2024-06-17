package controllers

import (
	"golang-template/dto"
	"golang-template/services"
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

type scannedNutritionController struct {
	SNService    services.ScannedNutritionService
	TokenService services.TokenService
}

func NewScannedNutritionController() *scannedNutritionController {
	return &scannedNutritionController{
		SNService:    services.NewScannedNutritionService(),
		TokenService: services.NewTokenService(),
	}
}

func (sn *scannedNutritionController) CreateScannedNutrition(c echo.Context) error {
	user, err := sn.TokenService.UserToken(c)
	if err != nil {
		response := map[string]interface{}{
			"code":    http.StatusForbidden,
			"status":  "failed",
			"message": "unauthorized",
		}
		return c.JSON(http.StatusForbidden, response)
	}

	err = sn.SNService.CreateScannedNutrition(c, user.ID)
	if err != nil {
		response := map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"status":  "failed",
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := map[string]interface{}{
		"code":    http.StatusOK,
		"status":  "success",
		"message": "Scanned nutrition created successfully",
	}
	return c.JSON(http.StatusOK, response)
}

func (sn *scannedNutritionController) GetScannedNutritionById(c echo.Context) error {
	id := c.Param("id")
	_, err := sn.TokenService.UserToken(c)
	if err != nil {
		response := map[string]interface{}{
			"code":    http.StatusForbidden,
			"status":  "failed",
			"message": "unauthorized",
		}
		return c.JSON(http.StatusForbidden, response)
	}

	scannedNutrition, err := sn.SNService.GetScannedNutritionById(id)
	if err != nil {
		response := map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"status":  "failed",
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := map[string]interface{}{
		"code":    http.StatusOK,
		"status":  "success",
		"message": "Scanned nutrition found",
		"data":    scannedNutrition,
	}
	return c.JSON(http.StatusOK, response)
}

func (sn *scannedNutritionController) GetScannedNutritionByUserId(c echo.Context) error {
	user, err := sn.TokenService.UserToken(c)
	if err != nil {
		response := map[string]interface{}{
			"code":    http.StatusForbidden,
			"status":  "failed",
			"message": "unauthorized",
		}
		return c.JSON(http.StatusForbidden, response)
	}

	page := 1
	pageSize := 10

	if qp := c.QueryParam("page"); qp != "" {
		if p, err := strconv.Atoi(qp); err == nil {
			page = p
		}
	}

	if qp := c.QueryParam("pageSize"); qp != "" {
		if ps, err := strconv.Atoi(qp); err == nil {
			pageSize = ps
		}
	}

	sort := c.QueryParam("sort")
	if sort != "" && !dto.IsValidSortField(sort) {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid sort fields"})
	}

	var desc int
	if qp := c.QueryParam("desc"); qp != "" {
		if ds, err := strconv.Atoi(qp); err == nil {
			desc = ds
		}
	}

	var search string
	if sp := c.QueryParam("search"); sp != "" {
		search = sp
	}

	scannedNutrition, pagination, err := sn.SNService.GetScannedNutritionByUserId(desc, page, pageSize, search, sort, user.ID)
	if err != nil {
		response := map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"status":  "failed",
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := map[string]interface{}{
		"code":       http.StatusOK,
		"status":     "success",
		"datas":      scannedNutrition,
		"pagination": pagination,
	}

	return c.JSON(http.StatusOK, response)
}

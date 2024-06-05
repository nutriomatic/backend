package controllers

import (
	"golang-template/dto"
	"golang-template/services"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type ProductTypeController struct {
	ProductTypeService services.ProductTypeService
	TokenService       services.TokenService
}

func NewProductTypeController() *ProductTypeController {
	return &ProductTypeController{
		ProductTypeService: services.NewProductTypeService(),
		TokenService:       services.NewTokenService(),
	}
}

func (pt *ProductTypeController) CreateProductType(c echo.Context) error {
	if !pt.TokenService.IsAdmin(c) {
		response := map[string]interface{}{
			"status":  "failed",
			"message": "Unauthorized access",
		}

		return c.JSON(http.StatusUnauthorized, response)
	}

	PTForm := &dto.PTRegisterForm{}
	err := c.Bind(PTForm)
	if err != nil {
		response := map[string]interface{}{
			"status":  "error",
			"message": "All product type fields must be provided!",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	err = pt.ProductTypeService.CreatePT(PTForm)
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
		"message": "Product Type created successfully",
	}
	return c.JSON(http.StatusCreated, response)
}

func (pt *ProductTypeController) GetAllProductType(c echo.Context) error {
	if !pt.TokenService.IsAdmin(c) {
		response := map[string]interface{}{
			"status":  "failed",
			"message": "Unauthorized access",
		}

		return c.JSON(http.StatusUnauthorized, response)
	}

	productTypes, err := pt.ProductTypeService.GetAllProductType()
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
		"data":   productTypes,
	}
	return c.JSON(http.StatusOK, response)
}

func (pt *ProductTypeController) GetProductTypeById(c echo.Context) error {
	if !pt.TokenService.IsAdmin(c) {
		response := map[string]interface{}{
			"status":  "failed",
			"message": "Unauthorized access",
		}

		return c.JSON(http.StatusUnauthorized, response)
	}

	id := c.Param("id")
	productType, err := pt.ProductTypeService.GetProductTypeById(id)
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
		"data":   productType,
	}

	return c.JSON(http.StatusOK, response)
}

func (pt *ProductTypeController) DeleteProductType(c echo.Context) error {
	if !pt.TokenService.IsAdmin(c) {
		response := map[string]interface{}{
			"status":  "failed",
			"message": "Unauthorized access",
		}

		return c.JSON(http.StatusUnauthorized, response)
	}

	id := c.Param("id")
	err := pt.ProductTypeService.DeletePT(id)
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
		"message": "Product Type deletion was successful",
	}
	return c.JSON(http.StatusOK, response)
}

package controllers

import (
	"golang-template/dto"
	"golang-template/services"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type productController struct {
	ProductService services.ProductService
	TokenService   services.TokenService
}

func NewProductController() *productController {
	return &productController{
		ProductService: services.NewProductService(),
		TokenService:   services.NewTokenService(),
	}
}

func (s *productController) CreateProduct(c echo.Context) error {
	user, err := s.TokenService.UserToken(c)
	if err != nil {
		return echo.ErrUnauthorized
	} else {
		productForm := &dto.ProductRegisterForm{}
		err := c.Bind(productForm)

		if err != nil {
			response := map[string]interface{}{
				"status":  "error",
				"message": "All product fields must be provided!",
			}
			return c.JSON(http.StatusBadRequest, response)
		}

		store, err := services.NewStoreService().GetStoreByUserId(user.ID)
		if err != nil {
			response := map[string]interface{}{
				"status":  "error",
				"message": "store not found",
			}
			return c.JSON(http.StatusNotFound, response)
		}

		err = s.ProductService.CreateProduct(productForm, store)
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
	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "product registration was successful",
	}
	return c.JSON(http.StatusOK, response)
}

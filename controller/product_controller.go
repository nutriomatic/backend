package controllers

import (
	"golang-template/dto"
	"golang-template/services"
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

type productController struct {
	ProductService services.ProductService
	TokenService   services.TokenService
	StoreService   services.StoreService
}

func NewProductController() *productController {
	return &productController{
		ProductService: services.NewProductService(),
		TokenService:   services.NewTokenService(),
		StoreService:   services.NewStoreService(),
	}
}

func (pc *productController) CreateProduct(c echo.Context) error {
	_, err := pc.TokenService.UserByToken(c)
	if err != nil {
		return echo.ErrUnauthorized
	}

	err = pc.ProductService.CreateProduct(c)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Product created successfully",
	}
	return c.JSON(http.StatusOK, response)
}

func (pc *productController) GetProductById(c echo.Context) error {
	id := c.Param("id")
	product, err := pc.ProductService.GetProductById(id)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	response := map[string]interface{}{
		"status":  "success",
		"product": product,
	}

	return c.JSON(http.StatusOK, response)
}

func (pc *productController) GetProductByStoreId(c echo.Context) error {
	id := c.Param("id")
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

	products, pagination, err := pc.ProductService.GetProductByStoreId(id, desc, page, pageSize, search, sort)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	response := map[string]interface{}{
		"status":     "success",
		"products":   products,
		"pagination": pagination,
	}

	return c.JSON(http.StatusOK, response)
}

func (pc *productController) GetAllProduct(c echo.Context) error {
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

	products, pagination, err := pc.ProductService.GetAllProduct(desc, page, pageSize, search, sort)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	response := map[string]interface{}{
		"status":     "success",
		"products":   products,
		"pagination": pagination,
	}

	return c.JSON(http.StatusOK, response)
}

func (pc *productController) UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	err := pc.ProductService.CheckProductStore(id, c)
	if err != nil {
		response := map[string]interface{}{
			"status":  "failed",
			"message": "unauthorized",
		}
		return c.JSON(http.StatusForbidden, response)
	}

	err = pc.ProductService.UpdateProduct(c, id)
	if err != nil {
		response := map[string]interface{}{
			"status":  "failed",
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Product updated successfully",
	}
	return c.JSON(http.StatusOK, response)
}

func (pc *productController) DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	err := pc.ProductService.CheckProductStore(id, c)
	if err != nil {
		response := map[string]interface{}{
			"status":  "failed",
			"message": "unauthorized",
		}
		return c.JSON(http.StatusForbidden, response)
	}

	err = pc.ProductService.DeleteProduct(id)
	if err != nil {
		response := map[string]interface{}{
			"status":  "failed",
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Product deleted successfully",
	}
	return c.JSON(http.StatusOK, response)
}

func (pc *productController) AdvertiseProduct(c echo.Context) error {
	id := c.Param("id")
	err := pc.ProductService.CheckProductStore(id, c)
	if err != nil {
		response := map[string]interface{}{
			"status":  "failed",
			"message": "unauthorized",
		}
		return c.JSON(http.StatusForbidden, response)
	}

	err = pc.ProductService.AdvertiseProduct(c, id)
	if err != nil {
		response := map[string]interface{}{
			"status":  "failed",
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Product advertised successfully",
	}
	return c.JSON(http.StatusOK, response)
}

func (pc *productController) UnadvertiseProduct(c echo.Context) error {
	id := c.Param("id")
	err := pc.ProductService.CheckProductStore(id, c)
	if err != nil {
		response := map[string]interface{}{
			"status":  "failed",
			"message": "unauthorized",
		}
		return c.JSON(http.StatusForbidden, response)
	}

	err = pc.ProductService.UnadvertiseProduct(c, id)
	if err != nil {
		response := map[string]interface{}{
			"status":  "failed",
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Product unadvertised successfully",
	}
	return c.JSON(http.StatusOK, response)
}

func (pc *productController) GetAllProductAdvertisement(c echo.Context) error {
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

	products, pagination, err := pc.ProductService.GetAllProductAdvertisement(desc, page, pageSize, search, sort)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	response := map[string]interface{}{
		"status":     "success",
		"products":   products,
		"pagination": pagination,
	}

	return c.JSON(http.StatusOK, response)
}

func (pc *productController) GetAllProductAdvertisementByStoreId(c echo.Context) error {
	id := c.Param("id")
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

	products, pagination, err := pc.ProductService.GetAllProductAdvertisementByStoreId(id, desc, page, pageSize, search, sort)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	response := map[string]interface{}{
		"status":     "success",
		"products":   products,
		"pagination": pagination,
	}

	return c.JSON(http.StatusOK, response)
}

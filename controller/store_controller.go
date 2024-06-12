package controllers

import (
	"golang-template/dto"
	"golang-template/services"
	"net/http"
	"strconv"

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
				"code":    http.StatusBadRequest,
				"status":  "error",
				"message": "All store fields must be provided!",
			}
			return c.JSON(http.StatusBadRequest, response)
		}

		isStore, _ := s.StoreService.GetStoreByUserId(user.ID)
		if isStore != nil {
			response := map[string]interface{}{
				"code":    http.StatusBadRequest,
				"status":  "error",
				"message": "store already exists",
			}
			return c.JSON(http.StatusBadRequest, response)
		}

		err = s.StoreService.CreateStore(StoreForm, user)
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
				"message": "internal server error",
			}
			return c.JSON(http.StatusInternalServerError, response)
		}

		response := map[string]interface{}{
			"code":    http.StatusOK,
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
					"code":    httpError.Code,
					"status":  "error",
					"message": httpError.Message,
				}
				return c.JSON(httpError.Code, response)
			}
			response := map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"status":  "error",
				"message": "internal server error",
			}
			return c.JSON(http.StatusInternalServerError, response)
		}
		response := map[string]interface{}{
			"code":   http.StatusOK,
			"status": "success",
			"store":  store,
		}
		return c.JSON(http.StatusOK, response)
	}
}

func (s *storeController) UpdateStore(c echo.Context) error {
	updateForm := &dto.StoreRegisterForm{}
	err := c.Bind(updateForm)
	if err != nil {
		response := map[string]interface{}{
			"code":    http.StatusBadRequest,
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
				"code":    httpError.Code,
				"status":  "error",
				"message": httpError.Message,
			}
			return c.JSON(httpError.Code, response)
		}
		response := map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"status":  "error",
			"message": "internal server error",
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := map[string]interface{}{
		"code":    http.StatusOK,
		"status":  "success",
		"message": "successfully updated store",
	}
	return c.JSON(http.StatusOK, response)
}

func (s *storeController) DeleteStore(c echo.Context) error {
	err := s.StoreService.DeleteStore(c)
	if err != nil {
		response := map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"status":  "error",
			"message": "error in removing store",
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := map[string]interface{}{
		"code":    http.StatusOK,
		"status":  "success",
		"message": "store removed successfully",
	}
	return c.JSON(http.StatusOK, response)
}

func (s *storeController) GetAllStores(c echo.Context) error {
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

	stores, pagination, err := s.StoreService.GetAll(desc, page, pageSize, search, sort)
	if err != nil {
		response := map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"status":  "error",
			"message": "error in fetching stores",
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := map[string]interface{}{
		"status":     "success",
		"message":    "stores fetched successfully",
		"stores":     stores,
		"pagination": pagination,
	}
	return c.JSON(http.StatusOK, response)
}

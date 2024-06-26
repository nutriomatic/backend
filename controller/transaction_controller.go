package controllers

import (
	"golang-template/dto"
	"golang-template/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type transactionController struct {
	TransactionService services.TransactionService
	TokenService       services.TokenService
	StoreService       services.StoreService
}

func NewTransactionController() *transactionController {
	return &transactionController{
		TransactionService: services.NewTransactionService(),
		TokenService:       services.NewTokenService(),
		StoreService:       services.NewStoreService(),
	}
}

func (tc *transactionController) CreateTransaction(c echo.Context) error {
	err := tc.TransactionService.CreateTransaction(c)
	if err != nil {
		response := map[string]interface{}{
			"code":    500,
			"status":  "failed",
			"message": err.Error(),
		}
		return c.JSON(500, response)
	}

	response := map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "Transaction created successfully",
	}
	return c.JSON(200, response)
}

func (tc *transactionController) GetTransactionById(c echo.Context) error {
	id := c.Param("id")
	transaction, err := tc.TransactionService.GetTransactionById(id)
	if err != nil {
		response := map[string]interface{}{
			"code":    500,
			"status":  "failed",
			"message": err.Error(),
		}
		return c.JSON(500, response)
	}

	response := map[string]interface{}{
		"code":        200,
		"status":      "success",
		"transaction": transaction,
	}
	return c.JSON(200, response)
}

func (tc *transactionController) GetTransactionByStoreId(c echo.Context) error {
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

	var sort string
	s := c.QueryParam("sort")
	if sort != "" {
		sort = s
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

	transactions, pagination, err := tc.TransactionService.GetTransactionByStoreId(id, desc, page, pageSize, search, sort)
	if err != nil {
		response := map[string]interface{}{
			"code":       http.StatusInternalServerError,
			"status":     "failed",
			"message":    err.Error(),
			"pagination": pagination,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := map[string]interface{}{
		"code":         http.StatusOK,
		"status":       "success",
		"transactions": transactions,
		"pagination":   pagination,
	}

	return c.JSON(http.StatusOK, response)
}

func (tc *transactionController) GetAllTransaction(c echo.Context) error {
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

	var sort string
	s := c.QueryParam("sort")
	if s != "" {
		sort = s
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

	var status string
	if st := c.QueryParam("status"); st != "" {
		status = st
	}

	transactions, pagination, err := tc.TransactionService.GetAllTransaction(desc, page, pageSize, search, sort, status)
	if err != nil {
		response := map[string]interface{}{
			"code":       http.StatusInternalServerError,
			"status":     "failed",
			"message":    err.Error(),
			"pagination": pagination,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := map[string]interface{}{
		"code":         http.StatusOK,
		"status":       "success",
		"transactions": transactions,
		"pagination":   pagination,
	}
	return c.JSON(http.StatusOK, response)
}

func (tc *transactionController) UpdateStatusTransaction(c echo.Context) error {
	id := c.Param("id")
	updateStatus := &dto.StatusTransaction{}
	err := c.Bind(updateStatus)
	if err != nil {
		response := map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "failed",
			"message": "Invalid request",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	err = tc.TransactionService.UpdateStatusTransaction(updateStatus.Status, c, id)
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
		"message": "Transaction status updated successfully",
	}
	return c.JSON(http.StatusOK, response)
}

func (tc *transactionController) DeleteTransaction(c echo.Context) error {
	err := tc.TransactionService.DeleteTransaction(c)
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
		"message": "Transaction removed successfully",
	}
	return c.JSON(http.StatusOK, response)
}

func (tc *transactionController) UploadProofPayment(c echo.Context) error {
	err := tc.TransactionService.UploadProofPayment(c)
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
		"message": "Payment proof uploaded successfully",
	}
	return c.JSON(http.StatusOK, response)
}

func (tc *transactionController) Checkout(c echo.Context) error {
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

	var sort string
	s := c.QueryParam("sort")
	if s != "" {
		sort = s
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

	var status string
	if st := c.QueryParam("status"); st != "" {
		status = st
	}

	userToken, err := tc.TokenService.UserToken(c)
	if err != nil {
		response := map[string]interface{}{
			"code":    http.StatusUnauthorized,
			"status":  "failed",
			"message": err.Error(),
		}
		return c.JSON(http.StatusUnauthorized, response)
	}
	store, err := tc.StoreService.GetStoreByUserId(userToken.ID)
	if err != nil {
		response := map[string]interface{}{
			"code":    http.StatusUnauthorized,
			"status":  "failed",
			"message": err.Error(),
		}
		return c.JSON(http.StatusUnauthorized, response)
	}

	tsc, pagination, err := tc.TransactionService.FindAllNewTransactions(desc, page, pageSize, search, sort, status, store.STORE_ID)
	if err != nil {
		response := map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"status":  "failed",
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	log.Print("tsc: ", tsc)
	response := map[string]interface{}{
		"code":         http.StatusOK,
		"status":       "success",
		"pagination":   pagination,
		"transactions": tsc,
	}
	return c.JSON(http.StatusOK, response)
}

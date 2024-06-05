package services

import (
	"golang-template/dto"
	"golang-template/models"
	"golang-template/repository"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type ProductService interface {
	CreateProduct(registerReq *dto.ProductRegisterForm, store *models.Store) error
}

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductService() ProductService {
	return &productService{
		productRepo: repository.NewProductRepositoryGORM(),
	}
}

func (s *productService) CreateProduct(registerReq *dto.ProductRegisterForm, store *models.Store) error {
	productName := strings.ToLower(registerReq.ProductName)
	if productName == strings.ToLower(productName) {
		return echo.NewHTTPError(http.StatusConflict, "Product name already exists")
	}
	return s.productRepo.CreateProduct(registerReq, store)
}

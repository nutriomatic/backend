package services

import (
	"golang-template/dto"
	"golang-template/models"
	"golang-template/repository"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ProductTypeService interface {
	GetProductTypeIdByType(cat int64) (id string, err error)
	GetProductTypeById(id string) (*models.ProductType, error)
	CreatePT(pt *dto.PTRegisterForm) error
	DeletePT(id string) error
	GetAllProductType() ([]models.ProductType, error)
}

type productTypeService struct {
	ptRepo repository.ProductTypeRepository
}

func NewProductTypeService() ProductTypeService {
	return &productTypeService{
		ptRepo: repository.NewProductTypeRepositoryGORM(),
	}
}

func (ptRepo *productTypeService) CreatePT(pt *dto.PTRegisterForm) error {
	existingPT, _ := ptRepo.ptRepo.GetProductTypeIdByType(pt.PTType)
	if existingPT != "" {
		return echo.NewHTTPError(http.StatusConflict, "Product type already exists")
	}

	newPT := &models.ProductType{
		PT_ID:     uuid.New().String(),
		PT_TYPE:   pt.PTType,
		PT_Name:   pt.PTName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return ptRepo.ptRepo.CreatePT(newPT)
}

func (ptRepo *productTypeService) GetAllProductType() ([]models.ProductType, error) {
	return ptRepo.ptRepo.GetAllProductType()
}

func (ptRepo *productTypeService) GetProductTypeIdByType(cat int64) (id string, err error) {
	return ptRepo.ptRepo.GetProductTypeIdByType(cat)
}

func (ptRepo *productTypeService) GetProductTypeById(id string) (*models.ProductType, error) {
	return ptRepo.ptRepo.GetProductTypeById(id)
}

func (ptRepo *productTypeService) DeletePT(id string) error {
	return ptRepo.ptRepo.DeletePT(id)
}

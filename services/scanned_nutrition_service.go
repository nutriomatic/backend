package services

import (
	"golang-template/dto"
	"golang-template/models"
	"golang-template/repository"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ScannedNutritionService interface {
	CreateScannedNutrition(c echo.Context, user_id string) error
	GetScannedNutritionById(id string) (*models.ScannedNutrition, error)
	GetScannedNutritionByUserId(desc, page, pageSize int, search, sort, id string) ([]models.ScannedNutrition, *dto.Pagination, error)
}

type scannedNutritionService struct {
	snRepo   repository.ScannedNutritionRepository
	uploader *ClientUploader
}

func NewScannedNutritionService() ScannedNutritionService {
	return &scannedNutritionService{
		snRepo:   repository.NewScannedNutritionRepositoryGORM(),
		uploader: NewClientUploader(),
	}
}

func (s *scannedNutritionService) CreateScannedNutrition(c echo.Context, user_id string) error {
	sn_name := c.FormValue("sn_name")

	imagePath, err := s.uploader.ProcessImageScannedNutrition(c)
	if err != nil {
		return err
	}
	realImagePath := "https://storage.googleapis.com/nutrio-storage/" + imagePath
	sn := models.ScannedNutrition{
		SN_ID:           uuid.New().String(),
		SN_PRODUCTNAME:  sn_name,
		SN_PRODUCTTYPE:  "",
		SN_INFO:         "",
		SN_PICTURE:      realImagePath,
		SN_ENERGY:       0,
		SN_PROTEIN:      0,
		SN_FAT:          0,
		SN_CARBOHYDRATE: 0,
		SN_SUGAR:        0,
		SN_SALT:         0,
		SN_GRADE:        "",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		USER_ID:         user_id,
	}

	return s.snRepo.CreateScannedNutrition(&sn)
}

func (s *scannedNutritionService) GetScannedNutritionById(id string) (*models.ScannedNutrition, error) {
	return s.snRepo.GetScannedNutritionById(id)
}

func (s *scannedNutritionService) GetScannedNutritionByUserId(desc, page, pageSize int, search, sort, id string) ([]models.ScannedNutrition, *dto.Pagination, error) {
	return s.snRepo.GetScannedNutritionByUserId(desc, page, pageSize, search, sort, id)
}

package services

import (
	"golang-template/dto"
	"golang-template/models"
	"golang-template/repository"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type ScannedNutritionService interface {
	CreateScannedNutrition(c echo.Context, user_id string) error
	GetScannedNutritionById(id string) (*models.ScannedNutrition, error)
	GetScannedNutritionByUserId(desc, page, pageSize int, search, sort, grade, id string) ([]models.ScannedNutrition, *dto.Pagination, error)
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
	err = godotenv.Load(".env")
	if err != nil {
		return err
	}
	realImagePath := os.Getenv("IMAGE_PATH") + imagePath

	err = godotenv.Load(".env")
	if err != nil {
		return err
	}
	url := os.Getenv("PYTHON_API") + "/ocr"

	requestData := &dto.SNRequest{
		Url: realImagePath,
	}

	responseData, _ := SendRequest[dto.SNRequest, dto.SNResponse](url, *requestData)

	sn := models.ScannedNutrition{
		SN_ID:           uuid.New().String(),
		SN_PRODUCTNAME:  sn_name,
		SN_PRODUCTTYPE:  "",
		SN_INFO:         "",
		SN_PICTURE:      realImagePath,
		SN_ENERGY:       responseData.Energy,
		SN_PROTEIN:      responseData.Protein,
		SN_FAT:          responseData.Fat,
		SN_CARBOHYDRATE: responseData.Carbs,
		SN_SUGAR:        responseData.Sugar,
		SN_SALT:         responseData.Salt,
		SN_GRADE:        responseData.Grade,
		SN_SATURATEDFAT: responseData.SaturatedFat,
		SN_FIBER:        responseData.Fiber,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		USER_ID:         user_id,
	}

	return s.snRepo.CreateScannedNutrition(&sn)
}

func (s *scannedNutritionService) GetScannedNutritionById(id string) (*models.ScannedNutrition, error) {
	return s.snRepo.GetScannedNutritionById(id)
}

func (s *scannedNutritionService) GetScannedNutritionByUserId(desc, page, pageSize int, search, sort, grade, id string) ([]models.ScannedNutrition, *dto.Pagination, error) {
	return s.snRepo.GetScannedNutritionByUserId(desc, page, pageSize, search, sort, grade, id)
}

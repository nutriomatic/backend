package services

import (
	"errors"
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
	// Get form values
	sn_name := c.FormValue("sn_name")

	// Process uploaded image
	imagePath, err := s.uploader.ProcessImageScannedNutrition(c)
	if err != nil {
		return err
	}

	// Load environment variables
	err = godotenv.Load(".env")
	if err != nil {
		return errors.New("failed to load environment variables: " + err.Error())
	}

	// Construct full image path and API URL
	realImagePath := os.Getenv("IMAGE_PATH") + imagePath
	url := os.Getenv("PYTHON_API") + "/ocr"

	// Prepare request data
	requestData := &dto.SNRequest{
		Url: realImagePath,
	}

	// Send request and handle response
	responseData, err := SendRequest[dto.SNRequest, dto.SNResponse](url, *requestData)
	if err != nil {
		return errors.New("error sending request to OCR API: " + err.Error())
	}

	// Create ScannedNutrition object
	sn := models.ScannedNutrition{
		SN_ID:           uuid.New().String(),
		SN_PRODUCTNAME:  sn_name,
		SN_PRODUCTTYPE:  "", // You might want to populate this field if needed
		SN_INFO:         "", // Add more fields if needed
		SN_PICTURE:      realImagePath,
		SN_ENERGY:       responseData.NutritionFacts.Energy,
		SN_PROTEIN:      responseData.NutritionFacts.Protein,
		SN_FAT:          responseData.NutritionFacts.Fat,
		SN_CARBOHYDRATE: responseData.NutritionFacts.Carbs,
		SN_SUGAR:        responseData.NutritionFacts.Sugar,
		SN_SALT:         responseData.NutritionFacts.Sodium,
		SN_GRADE:        responseData.Grade,
		SN_SATURATEDFAT: responseData.NutritionFacts.SaturatedFat,
		SN_FIBER:        responseData.NutritionFacts.Fiber,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		USER_ID:         user_id,
	}

	// Store the scanned nutrition data
	err = s.snRepo.CreateScannedNutrition(&sn)
	if err != nil {
		return errors.New("failed to store scanned nutrition data: " + err.Error())
	}

	return nil
}

func (s *scannedNutritionService) GetScannedNutritionById(id string) (*models.ScannedNutrition, error) {
	return s.snRepo.GetScannedNutritionById(id)
}

func (s *scannedNutritionService) GetScannedNutritionByUserId(desc, page, pageSize int, search, sort, grade, id string) ([]models.ScannedNutrition, *dto.Pagination, error) {
	return s.snRepo.GetScannedNutritionByUserId(desc, page, pageSize, search, sort, grade, id)
}

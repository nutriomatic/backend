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

type ActivityLevelService interface {
	GetActivityLevelIdByType(cat int64) (id string, err error)
	GetActivityLevelById(id string) (*models.ActivityLevel, error)
	CreateAL(al *dto.ALRegisterForm) error
	DeleteAL(id string) error
	GetAllActivityLevel() ([]models.ActivityLevel, error)
}

type activityLevelService struct {
	alRepo repository.ActivityLevelRepository
}

func NewActivityLevelService() ActivityLevelService {
	return &activityLevelService{
		alRepo: repository.NewActivityLevelRepositoryGORM(),
	}
}

func (alRepo *activityLevelService) CreateAL(al *dto.ALRegisterForm) error {
	existingAL, _ := alRepo.alRepo.GetActivityLevelIdByType(al.ALType)
	if existingAL != "" {
		return echo.NewHTTPError(http.StatusConflict, "Activity level type already exists")
	}

	newAL := &models.ActivityLevel{
		AL_ID:     uuid.New().String(),
		AL_TYPE:   al.ALType,
		AL_DESC:   al.ALDesc,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return alRepo.alRepo.CreateAL(newAL)
}

func (alRepo *activityLevelService) GetAllActivityLevel() ([]models.ActivityLevel, error) {
	return alRepo.alRepo.GetAllActivityLevel()
}

func (alRepo *activityLevelService) GetActivityLevelIdByType(cat int64) (id string, err error) {
	return alRepo.alRepo.GetActivityLevelIdByType(cat)
}

func (alRepo *activityLevelService) GetActivityLevelById(id string) (*models.ActivityLevel, error) {
	return alRepo.alRepo.GetActivityLevelById(id)
}

func (alRepo *activityLevelService) DeleteAL(id string) error {
	return alRepo.alRepo.DeleteAL(id)
}

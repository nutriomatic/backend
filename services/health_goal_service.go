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

type HealthGoalService interface {
	GetIdByType(cat int64) (id string, err error)
	GetById(id string) (*models.HealthGoal, error)
	CreateHealthGoal(hg *dto.HGRegisterForm) error
	DeleteHealthGoal(id string) error
	GetAllHealthGoal() ([]models.HealthGoal, error)
}

type healthGoalService struct {
	hgRepo repository.HealthGoalRepository
}

func NewHealthGoalService() HealthGoalService {
	return &healthGoalService{
		hgRepo: repository.NewHealthGoalRepositoryGORM(),
	}
}

func (hgRepo *healthGoalService) GetIdByType(cat int64) (id string, err error) {
	return hgRepo.hgRepo.GetIdByType(cat)
}

func (hgRepo *healthGoalService) GetById(id string) (*models.HealthGoal, error) {
	return hgRepo.hgRepo.GetById(id)
}

func (hgRepo *healthGoalService) CreateHealthGoal(hg *dto.HGRegisterForm) error {
	existingHG, _ := hgRepo.hgRepo.GetIdByType(hg.HGType)
	if existingHG != "" {
		return echo.NewHTTPError(http.StatusConflict, "Health goal type already exists")
	}

	newHG := &models.HealthGoal{
		HG_ID:     uuid.New().String(),
		HG_TYPE:   hg.HGType,
		HG_DESC:   hg.HGDesc,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return hgRepo.hgRepo.CreateHealthGoal(newHG)
}

func (hgRepo *healthGoalService) DeleteHealthGoal(id string) error {
	existingHG, _ := hgRepo.hgRepo.GetById(id)
	if existingHG == nil {
		return echo.NewHTTPError(http.StatusNotFound, "Health goal not found")
	}

	return hgRepo.hgRepo.DeleteHealthGoal(existingHG.HG_ID)
}

func (hgRepo *healthGoalService) GetAllHealthGoal() ([]models.HealthGoal, error) {
	return hgRepo.hgRepo.GetAllHealthGoal()
}

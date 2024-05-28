package services

import (
	"golang-template/models"
	"golang-template/repository"
)

type HealthGoalService interface {
	GetIdByType(cat int64) (id string, err error)
	GetById(id string) (*models.HealthGoal, error)
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

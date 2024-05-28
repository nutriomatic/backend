package services

import (
	"golang-template/models"
	"golang-template/repository"
)

type ActivityLevelService interface {
	GetIdByType(cat int64) (id string, err error)
	GetById(id string) (*models.ActivityLevel, error)
}

type activityLevelService struct {
	alRepo repository.ActivityLevelRepository
}

func NewActivityLevelService() ActivityLevelService {
	return &activityLevelService{
		alRepo: repository.NewActivityLevelRepositoryGORM(),
	}
}

func (alRepo *activityLevelService) GetIdByType(cat int64) (id string, err error) {
	return alRepo.alRepo.GetIdByType(cat)
}

func (alRepo *activityLevelService) GetById(id string) (*models.ActivityLevel, error) {
	return alRepo.alRepo.GetById(id)
}

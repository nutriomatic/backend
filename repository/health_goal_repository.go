package repository

import (
	"golang-template/config"
	"golang-template/models"

	"gorm.io/gorm"
)

type HealthGoalRepository interface {
	GetIdByType(cat int64) (id string, err error)
	GetById(id string) (*models.HealthGoal, error)
}

type HealthGoalRepositoryGORM struct {
	db *gorm.DB
}

func NewHealthGoalRepositoryGORM() *HealthGoalRepositoryGORM {
	return &HealthGoalRepositoryGORM{
		db: config.InitDB(),
	}
}

func (repo *HealthGoalRepositoryGORM) GetIdByType(cat int64) (id string, err error) {
	hg := &models.HealthGoal{}
	err = repo.db.Where("hg_type = ?", cat).First(hg).Error
	if err != nil {
		return "", err
	}
	return hg.HG_ID, nil
}

func (repo *HealthGoalRepositoryGORM) GetById(id string) (*models.HealthGoal, error) {
	var hg models.HealthGoal
	err := repo.db.Where("hg_id = ?", id).First(&hg).Error
	if err != nil {
		return nil, err
	}
	return &hg, nil
}

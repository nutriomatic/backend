package repository

import (
	"golang-template/config"
	"golang-template/models"

	"gorm.io/gorm"
)

type ActivityLevelRepository interface {
	GetIdByType(cat int64) (id string, err error)
	GetById(id string) (*models.ActivityLevel, error)
}

type ActivityLevelRepositoryGORM struct {
	db *gorm.DB
}

func NewActivityLevelRepositoryGORM() *ActivityLevelRepositoryGORM {
	return &ActivityLevelRepositoryGORM{
		db: config.InitDB(),
	}
}

func (repo *ActivityLevelRepositoryGORM) GetIdByType(cat int64) (id string, err error) {
	al := &models.ActivityLevel{}
	err = repo.db.Where("al_type = ?", cat).First(al).Error
	if err != nil {
		return "", err
	}
	return al.AL_ID, nil
}

func (repo *ActivityLevelRepositoryGORM) GetById(id string) (*models.ActivityLevel, error) {
	var al models.ActivityLevel
	err := repo.db.Where("al_id = ?", id).First(&al).Error
	if err != nil {
		return nil, err
	}
	return &al, nil
}

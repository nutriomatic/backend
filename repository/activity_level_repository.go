package repository

import (
	"golang-template/config"
	"golang-template/models"

	"gorm.io/gorm"
)

type ActivityLevelRepository interface {
	CreateAL(al *models.ActivityLevel) error
	GetIdByType(cat int64) (id string, err error)
	GetById(id string) (*models.ActivityLevel, error)
	GetAllActivityLevel() ([]models.ActivityLevel, error)
	UpdateAL(al *models.ActivityLevel) error
	DeleteAL(id string) error
}

type ActivityLevelRepositoryGORM struct {
	db *gorm.DB
}

func NewActivityLevelRepositoryGORM() *ActivityLevelRepositoryGORM {
	return &ActivityLevelRepositoryGORM{
		db: config.InitDB(),
	}
}

func (repo *ActivityLevelRepositoryGORM) CreateAL(al *models.ActivityLevel) error {
	err := repo.db.Create(al).Error
	if err != nil {
		return err
	}
	return nil
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

func (repo *ActivityLevelRepositoryGORM) GetAllActivityLevel() ([]models.ActivityLevel, error) {
	var al []models.ActivityLevel
	err := repo.db.Find(&al).Error
	if err != nil {
		return nil, err
	}
	return al, nil
}

func (repo *ActivityLevelRepositoryGORM) UpdateAL(al *models.ActivityLevel) error {
	err := repo.db.Save(al).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *ActivityLevelRepositoryGORM) DeleteAL(id string) error {
	err := repo.db.Where("al_id = ?", id).Delete(&models.ActivityLevel{}).Error
	if err != nil {
		return err
	}
	return nil
}

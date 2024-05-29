package repository

import (
	"golang-template/config"

	"golang-template/models"

	"gorm.io/gorm"
)

type StoreRepository interface {
	CreateStore(store *models.Store) error
	GetStoreById(id string) (*models.Store, error)
	GetStoreByUsername(username string) (*models.Store, error)
	UpdateStore(store *models.Store) error
	DeleteStore(id string) error
}

type StoreRepositoryGORM struct {
	db *gorm.DB
}

func NewStoreRepositoryGORM() *StoreRepositoryGORM {
	return &StoreRepositoryGORM{
		db: config.InitDB(),
	}
}

func (repo *StoreRepositoryGORM) CreateStore(store *models.Store) error {
	return repo.db.Create(store).Error
}

func (repo *StoreRepositoryGORM) GetStoreById(id string) (*models.Store, error) {
	store := &models.Store{}
	err := repo.db.Where("store_id = ?", id).First(store).Error
	if err != nil {
		return nil, err
	}
	return store, nil
}

func (repo *StoreRepositoryGORM) GetStoreByUsername(username string) (*models.Store, error) {
	var store models.Store
	err := repo.db.Where("store_username = ?", username).First(&store).Error
	if err != nil {
		return nil, err
	}
	return &store, nil
}

func (repo *StoreRepositoryGORM) UpdateStore(store *models.Store) error {
	return repo.db.Save(store).Error
}

func (repo *StoreRepositoryGORM) DeleteStore(id string) error {
	return repo.db.Where("store_id = ?", id).Delete(&models.Store{}).Error
}

package repository

import (
	"golang-template/config"
	"golang-template/dto"

	"golang-template/models"

	"gorm.io/gorm"
)

type StoreRepository interface {
	CreateStore(store *models.Store) error
	GetStoreByUserId(id string) (*models.Store, error)
	GetStoreByUsername(username string) (*models.Store, error)
	UpdateStore(store *models.Store) error
	DeleteStore(id string) error
	GetAll(desc, page, pageSize int, search, sort string) (*[]models.Store, *dto.Pagination, error)
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

func (repo *StoreRepositoryGORM) GetStoreByUserId(id string) (*models.Store, error) {
	store := &models.Store{}
	err := repo.db.Where("user_id = ?", id).First(store).Error
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

func (repo *StoreRepositoryGORM) GetAll(desc, page, pageSize int, search, sort string) (*[]models.Store, *dto.Pagination, error) {
	var store []models.Store
	query := repo.db.Model(&store)

	if search != "" {
		query = repo.db.Where("store_name ILIKE ? OR store_username ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if sort != "" {
		order := "ASC"
		if desc == 1 {
			order = "DESC"
		}
		query = query.Order(sort + " " + order)
	}

	pagination, err := dto.GetPaginated(query, page, pageSize, &store)
	if err != nil {
		return nil, nil, err
	}

	return &store, pagination, nil
}

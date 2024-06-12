package repository

import (
	"golang-template/config"
	"golang-template/dto"
	"golang-template/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(t *models.Transaction) error
	GetTransactionById(id string) (*models.Transaction, error)
	GetTransactionByStoreId(id string, desc, page, pageSize int, search, sort string) (*[]models.Transaction, *dto.Pagination, error)
	GetAllTransaction(desc, page, pageSize int, search, sort string) (*[]models.Transaction, *dto.Pagination, error)
	GetTransactionByUserId(id string, desc, page, pageSize int, search, sort string) (*[]models.Transaction, *dto.Pagination, error)
	UpdateTransaction(t *models.Transaction) error
	UpdateStatusTransaction(id string, status string) error
	DeleteTransaction(id string) error
}

type TransactionRepositoryGORM struct {
	db *gorm.DB
}

func NewTransactionProductRepositoryGORM() *TransactionRepositoryGORM {
	return &TransactionRepositoryGORM{
		db: config.InitDB(),
	}
}

func (repo *TransactionRepositoryGORM) CreateTransaction(t *models.Transaction) error {
	err := repo.db.Create(&t).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *TransactionRepositoryGORM) GetTransactionById(id string) (*models.Transaction, error) {
	var t models.Transaction
	err := repo.db.Where("tsc_id = ?", id).First(&t).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (repo *TransactionRepositoryGORM) GetTransactionByStoreId(id string, desc, page, pageSize int, search, sort string) (*[]models.Transaction, *dto.Pagination, error) {
	var t []models.Transaction
	query := repo.db.Where("store_id = ?", id).Find(&t)

	if search != "" {
		query = query.Where("tsc_status LIKE ?", "%"+search+"%")
	}

	if sort != "" {
		query = query.Order(sort)
	}

	pagination, err := dto.GetPaginated(query, page, pageSize, &t)
	if err != nil {
		return nil, nil, err
	}

	return &t, pagination, nil
}

func (repo *TransactionRepositoryGORM) GetTransactionByUserId(id string, desc, page, pageSize int, search, sort string) (*[]models.Transaction, *dto.Pagination, error) {
	var t []models.Transaction
	query := repo.db.Where("user_id = ?", id).Find(&t)

	if search != "" {
		query = query.Where("tsc_status LIKE ?", "%"+search+"%")
	}

	if sort != "" {
		query = query.Order(sort)
	}

	pagination, err := dto.GetPaginated(query, page, pageSize, &t)
	if err != nil {
		return nil, nil, err
	}

	return &t, pagination, nil
}

func (repo *TransactionRepositoryGORM) GetAllTransaction(desc, page, pageSize int, search, sort string) (*[]models.Transaction, *dto.Pagination, error) {
	var t []models.Transaction
	query := repo.db.Find(&t)

	if search != "" {
		query = query.Where("tsc_status LIKE ?", "%"+search+"%")
	}

	if sort != "" {
		query = query.Order(sort)
	}

	pagination, err := dto.GetPaginated(query, page, pageSize, &t)
	if err != nil {
		return nil, nil, err
	}

	return &t, pagination, nil
}

func (repo *TransactionRepositoryGORM) UpdateTransaction(t *models.Transaction) error {
	err := repo.db.Save(&t).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *TransactionRepositoryGORM) UpdateStatusTransaction(id string, status string) error {
	var t models.Transaction
	err := repo.db.Model(&t).Where("tsc_id = ?", id).Update("tsc_status", status).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *TransactionRepositoryGORM) DeleteTransaction(id string) error {
	var t models.Transaction
	err := repo.db.Where("tsc_id = ?", id).Delete(&t).Error
	if err != nil {
		return err
	}

	return nil
}

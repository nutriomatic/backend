package repository

import (
	"errors"
	"golang-template/config"
	"golang-template/dto"
	"golang-template/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(t *models.Transaction) error
	GetTransactionById(id string) (*models.Transaction, error)
	GetTransactionByStoreId(id string, desc, page, pageSize int, search, sort string) (*[]models.Transaction, *dto.Pagination, error)
	GetAllTransaction(desc, page, pageSize int, search, sort, status string) (*[]models.Transaction, *dto.Pagination, error)
	GetTransactionByUserId(id string, desc, page, pageSize int, search, sort string) (*[]models.Transaction, *dto.Pagination, error)
	UpdateTransaction(t *models.Transaction) error
	UpdateStatusTransaction(id string, status string) error
	DeleteTransaction(id string) error
	FindAllPendingByStoreId(storeId string) (*[]models.Transaction, error)
	FindAllNewTransactions(id string) (*[]models.Transaction, error)
}

type TransactionRepositoryGORM struct {
	db *gorm.DB
}

func NewTransactionRepositoryGORM() *TransactionRepositoryGORM {
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

func (repo *TransactionRepositoryGORM) GetAllTransaction(desc, page, pageSize int, search, sort, status string) (*[]models.Transaction, *dto.Pagination, error) {
	var t []models.Transaction
	query := repo.db.Where("tsc_bukti IS NOT NULL").Find(&t)

	if status != "" {
		query = query.Where("tsc_status LIKE ?", "%"+status+"%")
	}

	if sort != "" {
		order := "ASC"
		if desc == 1 {
			order = "DESC"
		}
		query = query.Order(sort + " " + order)
	}

	pagination, err := dto.GetPaginated(query, page, pageSize, &t)
	if err != nil {
		return nil, nil, err
	}

	return &t, pagination, nil
}

func (repo *TransactionRepositoryGORM) UpdateTransaction(transaction *models.Transaction) error {
	if transaction.TSC_ID == "" {
		return errors.New("transaction ID is required")
	}

	result := repo.db.Model(&models.Transaction{}).Where("tsc_id = ?", transaction.TSC_ID).Updates(transaction)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no rows were updated")
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

func (repo *TransactionRepositoryGORM) FindAllPendingByStoreId(storeId string) (*[]models.Transaction, error) {
	var t []models.Transaction
	err := repo.db.Where("store_id = ? AND tsc_status = 'pending'", storeId).Find(&t).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (repo *TransactionRepositoryGORM) FindAllNewTransactions(id string) (*[]models.Transaction, error) {
	var t []models.Transaction
	var p []models.Product

	err := repo.db.Where("product_isshow = 2").Find(&p).Error
	if err != nil {
		return nil, err
	}

	productIDs := make([]string, len(p))
	for i, product := range p {
		productIDs[i] = product.PRODUCT_ID
	}

	err = repo.db.Where("product_id IN (?)", productIDs).
		Where("tsc_bukti IS NULL").
		Where("store_id = ?", id).
		Find(&t).Error
	if err != nil {
		return nil, err
	}

	return &t, nil
}

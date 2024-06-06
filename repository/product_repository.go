package repository

import (
	"golang-template/config"
	"golang-template/dto"
	"golang-template/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(p *models.Product) error
	GetProductById(id string) (*models.Product, error)
	GetProductByStoreId(id string, desc, page, pageSize int, search, sort string) (*[]models.Product, *dto.Pagination, error)
	GetAllProduct(desc, page, pageSize int, search, sort string) (*[]models.Product, *dto.Pagination, error)
	UpdateProduct(p *models.Product) error
	DeleteProduct(id string) error
}

type ProductRepositoryGORM struct {
	db *gorm.DB
}

func NewProductRepositoryGORM() *ProductRepositoryGORM {
	return &ProductRepositoryGORM{
		db: config.InitDB(),
	}
}

func (repo *ProductRepositoryGORM) CreateProduct(p *models.Product) error {
	err := repo.db.Create(p).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *ProductRepositoryGORM) GetProductById(id string) (*models.Product, error) {
	var p models.Product
	err := repo.db.Where("product_id = ?", id).First(&p).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (repo *ProductRepositoryGORM) GetProductByStoreId(id string, desc, page, pageSize int, search, sort string) (*[]models.Product, *dto.Pagination, error) {
	var p []models.Product
	query := repo.db.Where("store_id = ?", id).Find(&p)

	if search != "" {
		query = query.Where("product_name LIKE ?", "%"+search+"%")
	}

	if sort != "" {
		query = query.Order(sort)
	}

	pagination, err := dto.GetPaginated(query, page, pageSize, &p)
	if err != nil {
		return nil, nil, err
	}
	return &p, pagination, nil
}

func (repo *ProductRepositoryGORM) GetAllProduct(desc, page, pageSize int, search, sort string) (*[]models.Product, *dto.Pagination, error) {
	var p []models.Product
	query := repo.db.Find(&p)

	if search != "" {
		query = query.Where("product_name LIKE ?", "%"+search+"%")
	}

	if sort != "" {
		query = query.Order(sort)
	}

	pagination, err := dto.GetPaginated(query, page, pageSize, &p)
	if err != nil {
		return nil, nil, err
	}
	return &p, pagination, nil
}

func (repo *ProductRepositoryGORM) UpdateProduct(p *models.Product) error {
	return repo.db.Save(p).Error
}

func (repo *ProductRepositoryGORM) DeleteProduct(id string) error {
	return repo.db.Where("product_id = ?", id).Delete(&models.Product{}).Error
}

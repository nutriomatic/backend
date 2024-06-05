package repository

import (
	"golang-template/config"
	"golang-template/dto"
	"golang-template/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(product *dto.ProductRegisterForm, store *models.Store) error
	GetProductById(id string) (*models.Product, error)
	GetProductByName(name string) (*models.Product, error)
	GetProductByStoreId(id string) (*[]models.Product, error)
}

type ProductRepositoryGORM struct {
	db *gorm.DB
}

func NewProductRepositoryGORM() *ProductRepositoryGORM {
	return &ProductRepositoryGORM{
		db: config.InitDB(),
	}
}

func (repo *ProductRepositoryGORM) CreateProduct(product *dto.ProductRegisterForm, store *models.Store) error {
	newProduct := models.Product{
		PRODUCT_ID:          uuid.New().String(),
		PRODUCT_NAME:        product.ProductName,
		PRODUCT_PRICE:       product.ProductPrice,
		PRODUCT_DESC:        product.ProductDesc,
		PRODUCT_ISSHOW:      product.ProductIsShow,
		PRODUCT_LEMAKTOTAL:  product.ProductLemakTotal,
		PRODUCT_PROTEIN:     product.ProductProtein,
		PRODUCT_KARBOHIDRAT: product.ProductKarbohidrat,
		PRODUCT_GARAM:       product.ProductGaram,
		STORE_ID:            store.STORE_ID,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}
	return repo.db.Create(&newProduct).Error
}

func (repo *ProductRepositoryGORM) GetProductById(id string) (*models.Product, error) {
	var product models.Product
	err := repo.db.Where("PRODUCT_ID = ?", id).First(&product).Error
	return &product, err
}

func (repo *ProductRepositoryGORM) GetProductByName(name string) (*models.Product, error) {
	var product models.Product
	err := repo.db.Where("PRODUCT_NAME = ?", name).First(&product).Error
	return &product, err
}

func (repo *ProductRepositoryGORM) GetProductByStoreId(id string) (*[]models.Product, error) {
	var products []models.Product
	err := repo.db.Where("STORE_ID = ?", id).Find(&products).Error
	return &products, err
}

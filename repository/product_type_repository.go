package repository

import (
	"golang-template/config"
	"golang-template/models"

	"gorm.io/gorm"
)

type ProductTypeRepository interface {
	CreatePT(pt *models.ProductType) error
	GetProductTypeIdByType(cat int64) (id string, err error)
	GetProductTypeById(id string) (*models.ProductType, error)
	GetAllProductType() ([]models.ProductType, error)
	UpdatePT(pt *models.ProductType) error
	DeletePT(id string) error
}

type ProductTypeRepositoryGORM struct {
	db *gorm.DB
}

func NewProductTypeRepositoryGORM() *ProductTypeRepositoryGORM {
	return &ProductTypeRepositoryGORM{
		db: config.InitDB(),
	}
}

func (repo *ProductTypeRepositoryGORM) CreatePT(pt *models.ProductType) error {
	err := repo.db.Create(pt).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *ProductTypeRepositoryGORM) GetProductTypeIdByType(cat int64) (id string, err error) {
	pt := &models.ProductType{}
	err = repo.db.Where("pt_type = ?", cat).First(pt).Error
	if err != nil {
		return "", err
	}
	return pt.PT_ID, nil
}

func (repo *ProductTypeRepositoryGORM) GetProductTypeById(id string) (*models.ProductType, error) {
	var pt models.ProductType
	err := repo.db.Where("pt_id = ?", id).First(&pt).Error
	if err != nil {
		return nil, err
	}
	return &pt, nil
}

func (repo *ProductTypeRepositoryGORM) GetAllProductType() ([]models.ProductType, error) {
	var pt []models.ProductType
	err := repo.db.Find(&pt).Error
	if err != nil {
		return nil, err
	}
	return pt, nil
}

func (repo *ProductTypeRepositoryGORM) UpdatePT(pt *models.ProductType) error {
	err := repo.db.Save(pt).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *ProductTypeRepositoryGORM) DeletePT(id string) error {
	err := repo.db.Where("pt_id = ?", id).Delete(&models.ProductType{}).Error
	if err != nil {
		return err
	}
	return nil
}

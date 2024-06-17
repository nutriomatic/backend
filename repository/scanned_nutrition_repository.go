package repository

import (
	"golang-template/config"
	"golang-template/dto"
	"golang-template/models"

	"gorm.io/gorm"
)

type ScannedNutritionRepository interface {
	CreateScannedNutrition(sn *models.ScannedNutrition) error
	GetScannedNutritionById(id string) (*models.ScannedNutrition, error)
	GetScannedNutritionByUserId(desc, page, pageSize int, search, sort, id string) ([]models.ScannedNutrition, *dto.Pagination, error)
}

type ScannedNutritionRepositoryGORM struct {
	db *gorm.DB
}

func NewScannedNutritionRepositoryGORM() *ScannedNutritionRepositoryGORM {
	return &ScannedNutritionRepositoryGORM{
		db: config.InitDB(),
	}
}

func (repo *ScannedNutritionRepositoryGORM) CreateScannedNutrition(sn *models.ScannedNutrition) error {
	err := repo.db.Create(&sn).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *ScannedNutritionRepositoryGORM) GetScannedNutritionById(id string) (*models.ScannedNutrition, error) {
	var sn models.ScannedNutrition
	err := repo.db.Where("sn_id = ?", id).First(&sn).Error
	if err != nil {
		return nil, err
	}
	return &sn, nil
}

func (repo *ScannedNutritionRepositoryGORM) GetScannedNutritionByUserId(desc, page, pageSize int, search, sort, id string) ([]models.ScannedNutrition, *dto.Pagination, error) {
	var sn []models.ScannedNutrition
	query := repo.db.Where("user_id = ?", id).Find(&sn)

	if search != "" {
		query = query.Where("sn_productname LIKE ?", "%"+search+"%")
	}

	if sort != "" {
		query = query.Order(sort)
	}

	pagination, err := dto.GetPaginated(query, page, pageSize, &sn)
	if err != nil {
		return nil, nil, err
	}
	return sn, pagination, nil
}

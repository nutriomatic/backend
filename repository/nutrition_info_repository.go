package repository

import (
	"golang-template/config"

	"gorm.io/gorm"
)

type NutritionInfoRepository interface {
}

type NutritionInfoRepositoryGORM struct {
	db *gorm.DB
}

func NewNutritionInfoRepositoryGORM() *NutritionInfoRepositoryGORM {
	return &NutritionInfoRepositoryGORM{
		db: config.InitDB(),
	}
}

package repository

import (
	"golang-template/config"
	"golang-template/models"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	GetPaymentIdByMethod(method string) (string, error)
}

type PaymentRepositoryGORM struct {
	db *gorm.DB
}

func NewPaymentRepositoryGORM() *PaymentRepositoryGORM {
	return &PaymentRepositoryGORM{
		db: config.InitDB(),
	}
}

func (repo *PaymentRepositoryGORM) GetPaymentIdByMethod(method string) (string, error) {
	var p models.Payment
	err := repo.db.Where("payment_method = ?", method).First(&p).Error
	if err != nil {
		return "", err
	}
	return p.PAYMENT_ID, nil
}

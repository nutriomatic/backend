package repository

import (
	"golang-template/config"
	"golang-template/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TokenRepository interface {
	SaveToken(user *models.User, token string) error
	UserByToken(token string) (*models.User, error)
}

type TokenRepositoryGORM struct {
	db *gorm.DB
}

func NewTokenRepositoryGORM() *TokenRepositoryGORM {
	return &TokenRepositoryGORM{
		db: config.InitDB(),
	}
}

func (repo *TokenRepositoryGORM) SaveToken(user *models.User, token string) error {
	AccessToken := models.Token{ID: uuid.New(), UserId: user.ID, Token: token, ExpiresAt: time.Now().Add(time.Hour * 24)}
	result := repo.db.Create(&AccessToken)

	return result.Error
}

func (repo *TokenRepositoryGORM) UserByToken(token string) (*models.User, error) {
	var AccessToken models.Token
	var User models.User

	err := repo.db.Where("token = ?", token).Where("expires_at > ?", time.Now()).First(&AccessToken).Error
	if err != nil {
		return nil, err
	}
	err = repo.db.Select("id", "username", "name", "email", "role", "created_at", "updated_at").Where("id = ?", AccessToken.UserId).First(&User).Error
	if err != nil {
		return nil, err
	}
	return &User, nil
}

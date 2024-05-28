package repository

import (
	"golang-template/config"
	"golang-template/dto"
	"golang-template/models"
	"time"

	"gorm.io/gorm"
)

type TokenRepository interface {
	SaveToken(user *models.User, token string) error
	UserByToken(token string) (*dto.UserResponseToken, error)
}

type TokenRepositoryGORM struct {
	db       *gorm.DB
	userRepo UserRepositoryGORM
}

func NewTokenRepositoryGORM() *TokenRepositoryGORM {
	return &TokenRepositoryGORM{
		db:       config.InitDB(),
		userRepo: UserRepositoryGORM{},
	}
}

func (repo *TokenRepositoryGORM) SaveToken(user *models.User, token string) error {
	AccessToken := models.Token{UserId: user.ID, Token: token, ExpiresAt: time.Now().Add(time.Hour * 24)}
	result := repo.db.Create(&AccessToken)

	return result.Error
}

func (repo *TokenRepositoryGORM) UserByToken(token string) (*dto.UserResponseToken, error) {
	var AccessToken models.Token
	var User models.User

	err := repo.db.Select("user_id").Where("token = ?", token).Where("expires_at > ?", time.Now()).First(&AccessToken).Error
	if err != nil {
		return nil, err
	}

	err = repo.db.Select("id", "username", "name", "email", "role", "gender", "telp", "profpic", "birthdate", "place", "height", "weight", "weight_goal", "created_at", "updated_at", "hg_id", "al_id").
		Where("id = ?", AccessToken.UserId).First(&User).Error
	if err != nil {
		return nil, err
	}

	al, err := NewActivityLevelRepositoryGORM().GetById(User.AL_ID)
	if err != nil {
		return nil, err
	}

	hg, err := NewHealthGoalRepositoryGORM().GetById(User.HG_ID)
	if err != nil {
		return nil, err
	}

	response := &dto.UserResponseToken{
		Id:         User.ID,
		Name:       User.Name,
		Username:   User.Username,
		Email:      User.Email,
		Role:       User.Role,
		Gender:     User.Gender,
		Telp:       User.Telp,
		Profpic:    User.Profpic,
		Birthdate:  User.Birthdate,
		Place:      User.Place,
		Height:     User.Height,
		Weight:     User.Weight,
		WeightGoal: User.WeightGoal,
		HG_ID:      User.HG_ID,
		HG_TYPE:    hg.HG_TYPE,
		HG_DESC:    hg.HG_DESC,
		AL_ID:      User.AL_ID,
		AL_TYPE:    al.AL_TYPE,
		AL_DESC:    al.AL_DESC,
		AL_VALUE:   al.AL_VALUE,
	}

	return response, nil
}

package repository

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"golang-template/config"
	"golang-template/dto"
	"golang-template/models"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserById(id uuid.UUID) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(id uuid.UUID) error
	FindAll(page, pageSize int, search, sort string) ([]models.User, *dto.Pagination, error)
	Logout(token string) error
}

type UserRepositoryGORM struct {
	db *gorm.DB
}

func NewUserRepositoryGORM() *UserRepositoryGORM {
	return &UserRepositoryGORM{
		db: config.InitDB(),
	}
}

func (repo *UserRepositoryGORM) CreateUser(user *models.User) error {
	return repo.db.Create(user).Error
}

func (repo *UserRepositoryGORM) GetUserById(id uuid.UUID) (*models.User, error) {
	user := &models.User{}
	err := repo.db.Where("id = ?", id).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepositoryGORM) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := repo.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepositoryGORM) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := repo.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepositoryGORM) GetUserByRole(role string) (*models.User, error) {
	var user models.User
	err := repo.db.Where("role = ?", role).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepositoryGORM) UpdateUser(user *models.User) (*models.User, error) {
	return user, repo.db.Save(user).Error
}

func (repo *UserRepositoryGORM) DeleteUser(id uuid.UUID) error {
	return repo.db.Where("id = ?", id).Delete(models.User{}).Error
}

func (r *UserRepositoryGORM) FindAll(page, pageSize int, search, sort string) ([]models.User, *dto.Pagination, error) {
	var users []models.User

	allUser := r.db.Find(&users)
	if allUser.Error != nil {
		return nil, nil, allUser.Error
	}

	if search != "" {
		allUser = r.db.Where("username ILIKE ? OR name ILIKE ? OR email ILIKE ? OR role ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").Find(&users)
		if allUser.Error != nil {
			return nil, nil, allUser.Error
		}
	}

	if sort != "" {
		allUser = r.db.Order(sort).Find(&users)
		if allUser.Error != nil {
			return nil, nil, allUser.Error
		}
	}

	pagination, err := dto.GetPaginated(allUser, page, pageSize, &users)
	if err != nil {
		return nil, nil, err
	}

	return users, pagination, nil

}

func (repo *UserRepositoryGORM) Logout(token string) error {
	var AccessToken models.Token
	err := repo.db.Where("token = ?", token).Where("expires_at > ?", time.Now()).First(&AccessToken).Error
	if err != nil {
		return err
	}
	repo.db.Where("token = ?", token).Where("expires_at > ?", time.Now()).Delete(&AccessToken)

	return nil
}

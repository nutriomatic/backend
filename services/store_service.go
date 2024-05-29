package services

import (
	"golang-template/dto"
	"golang-template/models"
	"golang-template/repository"
	"time"

	"github.com/google/uuid"
)

type StoreService interface {
	CreateStore(registerReq *dto.StoreRegisterForm, user *models.User) error
}

type storeService struct {
	storeRepo repository.StoreRepository
}

func NewStoreService() StoreService {
	return &storeService{
		storeRepo: repository.NewStoreRepositoryGORM(),
	}
}

func (s *storeService) CreateStore(registerReq *dto.StoreRegisterForm, user *models.User) error {
	newStore := models.Store{
		STORE_ID:       uuid.New().String(),
		STORE_NAME:     registerReq.StoreName,
		STORE_USERNAME: registerReq.StoreUsername,
		STORE_ADDRESS:  registerReq.StoreAddress,
		STORE_CONTACT:  registerReq.StoreContact,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		USER_ID:        user.ID,
	}

	return s.storeRepo.CreateStore(&newStore)
}

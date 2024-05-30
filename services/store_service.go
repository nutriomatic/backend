package services

import (
	"golang-template/dto"
	"golang-template/middleware"
	"golang-template/models"
	"golang-template/repository"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type StoreService interface {
	CreateStore(registerReq *dto.StoreRegisterForm, user *models.User) error
	GetStoreByUserId(id string) (*models.Store, error)
	GetStoreByUsername(username string) (*models.Store, error)
	UpdateStore(updateForm *dto.StoreRegisterForm, c echo.Context) error
}

type storeService struct {
	storeRepo repository.StoreRepository
	tokenRepo repository.TokenRepository
}

func NewStoreService() StoreService {
	return &storeService{
		storeRepo: repository.NewStoreRepositoryGORM(),
		tokenRepo: repository.NewTokenRepositoryGORM(),
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

func (s *storeService) GetStoreByUserId(id string) (*models.Store, error) {
	return s.storeRepo.GetStoreByUserId(id)
}

func (s *storeService) GetStoreByUsername(username string) (*models.Store, error) {
	return s.storeRepo.GetStoreByUsername(username)
}

func (s *storeService) UpdateStore(updateForm *dto.StoreRegisterForm, c echo.Context) error {
	tokenUser, err := s.tokenRepo.UserToken(middleware.GetToken(c))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}

	existingUser, err := s.storeRepo.GetStoreByUserId(tokenUser.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Store not found")
	}

	if updateForm.StoreName != "" {
		existingUser.STORE_NAME = updateForm.StoreName
	}

	if updateForm.StoreUsername != "" {
		if _, err := s.storeRepo.GetStoreByUsername(updateForm.StoreUsername); err == nil {
			return echo.NewHTTPError(http.StatusConflict, "Username already exists")
		}
		existingUser.STORE_USERNAME = updateForm.StoreUsername
	}

	if updateForm.StoreAddress != "" {
		existingUser.STORE_ADDRESS = updateForm.StoreAddress
	}

	if updateForm.StoreContact != "" {
		existingUser.STORE_CONTACT = updateForm.StoreContact
	}

	existingUser.UpdatedAt = time.Now()
	return s.storeRepo.UpdateStore(existingUser)
}

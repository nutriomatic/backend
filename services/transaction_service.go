package services

import (
	"golang-template/dto"
	"golang-template/middleware"
	"golang-template/models"
	"golang-template/repository"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type TransactionService interface {
	CreateTransaction(c echo.Context) error
	GetTransactionById(id string) (*models.Transaction, error)
	GetTransactionByStoreId(id string, desc, page, pageSize int, search, sort string) (*[]models.Transaction, *dto.Pagination, error)
	GetAllTransaction(desc, page, pageSize int, search, sort string) (*[]models.Transaction, *dto.Pagination, error)
	GetTransactionByUserId(id string, desc, page, pageSize int, search, sort string) (*[]models.Transaction, *dto.Pagination, error)
	UpdateStatusTransaction(status string, c echo.Context, id string) error
	DeleteTransaction(c echo.Context, id string) error
	UploadProofPayment(c echo.Context) error
}

type transactionService struct {
	tscRepo     repository.TransactionRepository
	tokenRepo   repository.TokenRepository
	storeRepo   repository.StoreRepository
	productRepo repository.ProductRepository
	uploader    *ClientUploader
}

func NewTransactionService() TransactionService {
	return &transactionService{
		tscRepo:     repository.NewTransactionProductRepositoryGORM(),
		tokenRepo:   repository.NewTokenRepositoryGORM(),
		storeRepo:   repository.NewStoreRepositoryGORM(),
		productRepo: repository.NewProductRepositoryGORM(),
		uploader:    NewClientUploader(),
	}
}

func (s *transactionService) CreateTransaction(c echo.Context) error {
	id := c.Param("product_id")
	userToken, err := s.tokenRepo.UserToken(middleware.GetToken(c))
	if err != nil {
		return err
	}

	store, err := s.storeRepo.GetStoreByUserId(userToken.ID)
	if err != nil {
		return err
	}

	availStore, err := s.productRepo.GetStoreByProductId(id)
	if err != nil {
		return err
	}

	if store.STORE_ID != availStore.STORE_ID {
		return err
	}

	newTsc := &models.Transaction{
		TSC_ID:     uuid.New().String(),
		TSC_PRICE:  5000.00,
		TSC_START:  time.Now(),
		TSC_END:    time.Now().AddDate(0, 0, 1),
		TSC_STATUS: "pending",
		STORE_ID:   store.STORE_ID,
		PRODUCT_ID: id,
	}

	return s.tscRepo.CreateTransaction(newTsc)
}

func (s *transactionService) GetTransactionById(id string) (*models.Transaction, error) {
	return s.tscRepo.GetTransactionById(id)
}

func (s *transactionService) GetTransactionByStoreId(id string, desc, page, pageSize int, search, sort string) (*[]models.Transaction, *dto.Pagination, error) {
	return s.tscRepo.GetTransactionByStoreId(id, desc, page, pageSize, search, sort)
}

func (s *transactionService) GetAllTransaction(desc, page, pageSize int, search, sort string) (*[]models.Transaction, *dto.Pagination, error) {
	return s.tscRepo.GetAllTransaction(desc, page, pageSize, search, sort)
}

func (s *transactionService) GetTransactionByUserId(id string, desc, page, pageSize int, search, sort string) (*[]models.Transaction, *dto.Pagination, error) {
	return s.tscRepo.GetTransactionByUserId(id, desc, page, pageSize, search, sort)
}

func (s *transactionService) UpdateStatusTransaction(status string, c echo.Context, id string) error {
	userToken, err := s.tokenRepo.UserToken(middleware.GetToken(c))
	if err != nil {
		return err
	}

	store, err := s.storeRepo.GetStoreByUserId(userToken.ID)
	if err != nil {
		return err
	}

	availStore, err := s.productRepo.GetStoreByProductId(id)
	if err != nil {
		return err
	}

	if store.STORE_ID != availStore.STORE_ID {
		return err
	}

	tsc, err := s.tscRepo.GetTransactionById(id)
	if err != nil {
		return err
	}

	if tsc.STORE_ID != store.STORE_ID {
		return err
	}

	tsc.TSC_STATUS = status
	tsc.STORE.UpdatedAt = time.Now()

	return s.tscRepo.UpdateTransaction(tsc)
}

func (s *transactionService) DeleteTransaction(c echo.Context, id string) error {
	userToken, err := s.tokenRepo.UserToken(middleware.GetToken(c))
	if err != nil {
		return err
	}

	store, err := s.storeRepo.GetStoreByUserId(userToken.ID)
	if err != nil {
		return err
	}

	availStore, err := s.productRepo.GetStoreByProductId(id)
	if err != nil {
		return err
	}

	if store.STORE_ID != availStore.STORE_ID {
		return err
	}

	tsc, err := s.tscRepo.GetTransactionById(id)
	if err != nil {
		return err
	}

	if tsc.STORE_ID != store.STORE_ID {
		return err
	}

	return s.tscRepo.DeleteTransaction(id)
}

func (s *transactionService) UploadProofPayment(c echo.Context) error {
	id := c.Param("id")

	userToken, err := s.tokenRepo.UserToken(middleware.GetToken(c))
	if err != nil {
		return err
	}

	store, err := s.storeRepo.GetStoreByUserId(userToken.ID)
	if err != nil {
		return err
	}

	tsc, err := s.tscRepo.GetTransactionById(id)
	if err != nil {
		return err
	}

	if tsc.STORE_ID != store.STORE_ID {
		return err
	}

	imagePath, err := s.uploader.ProcessImageProof(c)
	if err != nil {
		return err
	}
	realImagePath := "https://storage.googleapis.com/nutrio-storage/" + imagePath

	tsc, err = s.tscRepo.GetTransactionById(id)
	if err != nil {
		return err
	}

	tsc.TSC_BUKTI = realImagePath
	tsc.UpdatedAt = time.Now()

	return s.tscRepo.UpdateTransaction(tsc)
}

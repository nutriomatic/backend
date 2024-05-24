package services

import (
	"golang-template/middleware"
	"golang-template/models"
	"golang-template/repository"

	"github.com/labstack/echo/v4"
)

type TokenService interface {
	SaveToken(user *models.User, token string) error
	UserByToken(c echo.Context) (*models.User, error)
}

type tokenService struct {
	tokenRepo repository.TokenRepository
}

func NewTokenService() TokenService {
	return &tokenService{
		tokenRepo: repository.NewTokenRepositoryGORM(),
	}
}

func (s *tokenService) SaveToken(user *models.User, token string) error {
	return s.tokenRepo.SaveToken(user, token)
}

func (s *tokenService) UserByToken(c echo.Context) (*models.User, error) {
	token := middleware.GetToken(c)
	return s.tokenRepo.UserByToken(token)
}

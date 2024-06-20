package services

import (
	"golang-template/dto"
	"golang-template/middleware"
	"golang-template/models"
	"golang-template/repository"

	"github.com/labstack/echo/v4"
)

type TokenService interface {
	SaveToken(user *models.User, token string) error
	UserByToken(c echo.Context) (*dto.UserResponseToken, error)
	UserToken(c echo.Context) (*models.User, error)
	IsAdmin(c echo.Context) bool
	CheckToken(c echo.Context, token string) error
}

type tokenService struct {
	tokenRepo repository.TokenRepository
}

func NewTokenService() TokenService {
	return &tokenService{
		tokenRepo: repository.NewTokenRepositoryGORM(),
	}
}

func (s *tokenService) CheckToken(c echo.Context, token string) error {
	status := s.tokenRepo.FindToken(token)
	if status {
		response := map[string]interface{}{
			"code":    200,
			"status":  "success",
			"message": "Token is valid",
		}
		return c.JSON(200, response)
	} else {
		response := map[string]interface{}{
			"code":    401,
			"status":  "error",
			"message": "Token is invalid",
		}
		return c.JSON(401, response)
	}
}

func (s *tokenService) SaveToken(user *models.User, token string) error {
	return s.tokenRepo.SaveToken(user, token)
}

func (s *tokenService) UserByToken(c echo.Context) (*dto.UserResponseToken, error) {
	token := middleware.GetToken(c)
	return s.tokenRepo.UserByToken(token)
}

func (s *tokenService) UserToken(c echo.Context) (*models.User, error) {
	token := middleware.GetToken(c)
	return s.tokenRepo.UserToken(token)
}

func (s *tokenService) IsAdmin(c echo.Context) bool {
	user, err := s.UserByToken(c)
	if err != nil {
		return false
	}

	if user.Role == "admin" {
		return true
	} else {
		return false
	}
}

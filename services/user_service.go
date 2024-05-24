package services

import (
	"golang-template/dto"
	"golang-template/middleware"
	"golang-template/models"
	"golang-template/repository"
	"golang-template/utils"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserService interface {
	CreateUser(registerReq *dto.RegisterForm) error
	GetUserById(id uuid.UUID) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	UpdateUser(updateForm *dto.RegisterForm, c echo.Context) (*models.User, error)
	DeleteUser(c echo.Context) error
	Logout(c echo.Context) error
}

type userService struct {
	userRepo  repository.UserRepository
	tokenRepo repository.TokenRepository
}

func NewUserService() UserService {
	return &userService{
		userRepo: repository.NewUserRepositoryGORM(),
	}
}

func (s *userService) CreateUser(registerReq *dto.RegisterForm) error {
	newUser := models.User{
		ID:        uuid.New(),
		Name:      registerReq.Name,
		Username:  registerReq.Username,
		Email:     registerReq.Email,
		Password:  registerReq.Password,
		Role:      "customer",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	hashed, err := utils.HashPassword(newUser.Password)
	if err != nil {
		return err
	}
	newUser.Password = hashed
	return s.userRepo.CreateUser(&newUser)
}

func (s *userService) GetUserByUsername(username string) (*models.User, error) {
	return s.userRepo.GetUserByUsername(username)
}

func (s *userService) GetUserByEmail(email string) (*models.User, error) {
	return s.userRepo.GetUserByEmail(email)
}

func (s *userService) GetUserById(id uuid.UUID) (*models.User, error) {
	return s.userRepo.GetUserById(id)
}

func (s *userService) UpdateUser(updateForm *dto.RegisterForm, c echo.Context) (*models.User, error) {
	tokenUser, err := s.tokenRepo.UserByToken(middleware.GetToken(c))
	if err != nil {
		return nil, c.String(http.StatusUnauthorized, "Unauthorized")
	}

	existingUser, err := s.userRepo.GetUserById(tokenUser.ID)
	if err != nil {
		return nil, c.String(http.StatusInternalServerError, "error retrieving user")
	}

	if updateForm.Name != "" {
		existingUser.Name = updateForm.Name
	}

	if updateForm.Username != "" {
		if _, err := s.userRepo.GetUserByUsername(updateForm.Username); err == nil {
			return nil, c.String(http.StatusBadRequest, "username already exists")
		}
		existingUser.Username = updateForm.Username
	}

	if updateForm.Email != "" {
		if _, err := s.userRepo.GetUserByEmail(updateForm.Email); err == nil {
			return nil, c.String(http.StatusBadRequest, "email already exists")
		}
		existingUser.Email = updateForm.Email
	}

	if updateForm.Password != "" {
		if !utils.ValidateLengthPassword(updateForm.Password) {
			return nil, c.String(http.StatusBadRequest, "invalid password format")
		}
		hashedPassword, err := utils.HashPassword(updateForm.Password)
		if err != nil {
			return nil, c.String(http.StatusInternalServerError, "error hashing password")
		}
		existingUser.Password = hashedPassword
	}
	return s.userRepo.UpdateUser(existingUser)
}

func (s *userService) DeleteUser(c echo.Context) error {
	tokenUser, err := s.tokenRepo.UserByToken(middleware.GetToken(c))
	if err != nil {
		return c.String(http.StatusUnauthorized, "Unauthorized")
	}

	existingUser, err := s.userRepo.GetUserById(tokenUser.ID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "error retrieving user")
	}
	return s.userRepo.DeleteUser(existingUser.ID)
}

func (s *userService) Logout(c echo.Context) error {
	token := middleware.GetToken(c)
	return s.userRepo.Logout(token)
}

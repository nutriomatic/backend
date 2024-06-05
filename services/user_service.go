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
	CreateUser(registerReq *dto.Register) error
	GetUserById(id string) (*models.User, error)
	// GetUserByUsername(username string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	UpdateUser(updateForm *dto.RegisterForm, c echo.Context) (*dto.UserResponseToken, error)
	DeleteUser(c echo.Context) error
	Logout(c echo.Context) error
}

type userService struct {
	userRepo  repository.UserRepository
	tokenRepo repository.TokenRepository
}

func NewUserService() UserService {
	return &userService{
		userRepo:  repository.NewUserRepositoryGORM(),
		tokenRepo: repository.NewTokenRepositoryGORM(),
	}
}

func (s *userService) CreateUser(registerReq *dto.Register) error {
	al_id, err := NewActivityLevelService().GetActivityLevelIdByType(registerReq.AL_TYPE)
	if err != nil {
		return err
	}
	hg_id, err := NewHealthGoalService().GetIdByType(registerReq.HG_TYPE)
	if err != nil {
		return err
	}

	newUser := models.User{
		ID:        uuid.New().String(),
		Name:      registerReq.Name,
		Email:     registerReq.Email,
		Password:  registerReq.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		HG_ID:     hg_id,
		AL_ID:     al_id,
	}

	hashed, err := utils.HashPassword(newUser.Password)
	if err != nil {
		return err
	}
	newUser.Password = hashed
	return s.userRepo.CreateUser(&newUser)
}

// func (s *userService) GetUserByUsername(username string) (*models.User, error) {
// 	return s.userRepo.GetUserByUsername(username)
// }

func (s *userService) GetUserByEmail(email string) (*models.User, error) {
	return s.userRepo.GetUserByEmail(email)
}

func (s *userService) GetUserById(id string) (*models.User, error) {
	return s.userRepo.GetUserById(id)
}

func (s *userService) UpdateUser(updateForm *dto.RegisterForm, c echo.Context) (*dto.UserResponseToken, error) {
	tokenUser, err := s.tokenRepo.UserToken(middleware.GetToken(c))
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}

	existingUser, err := s.userRepo.GetUserById(tokenUser.ID)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "error retrieving user")
	}

	if updateForm.Name != "" {
		existingUser.Name = updateForm.Name
	}

	// if updateForm.Username != "" {
	// 	if _, err := s.userRepo.GetUserByUsername(updateForm.Username); err == nil {
	// 		return nil, echo.NewHTTPError(http.StatusBadRequest, "username already exists")
	// 	}
	// 	existingUser.Username = updateForm.Username
	// }

	if updateForm.Email != "" {
		if _, err := s.userRepo.GetUserByEmail(updateForm.Email); err == nil {
			return nil, echo.NewHTTPError(http.StatusBadRequest, "email already exists")
		}
		existingUser.Email = updateForm.Email
	}

	if updateForm.Password != "" {
		if !utils.ValidateLengthPassword(updateForm.Password) {
			return nil, echo.NewHTTPError(http.StatusBadRequest, "invalid password format")
		}
		hashedPassword, err := utils.HashPassword(updateForm.Password)
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, "error hashing password")
		}
		existingUser.Password = hashedPassword
	}

	if updateForm.Gender != 0 {
		existingUser.Gender = updateForm.Gender
	}

	if updateForm.Telp != "" {
		existingUser.Telp = updateForm.Telp
	}

	if updateForm.Profpic != "" {
		existingUser.Profpic = updateForm.Profpic
	}

	if updateForm.Birthdate != "" {
		existingUser.Birthdate = updateForm.Birthdate
	}

	if updateForm.Place != "" {
		existingUser.Place = updateForm.Place
	}

	if updateForm.Height != 0 {
		existingUser.Height = updateForm.Height
	}

	if updateForm.Weight != 0 {
		existingUser.Weight = updateForm.Weight
	}

	if updateForm.WeightGoal != 0 {
		existingUser.WeightGoal = updateForm.WeightGoal
	}

	if updateForm.AL_TYPE != 0 {
		al_id, err := NewActivityLevelService().GetActivityLevelIdByType(updateForm.AL_TYPE)
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, "error retrieving activity level")
		}
		existingUser.AL_ID = al_id
	}

	if updateForm.HG_TYPE != 0 {
		hg_id, err := NewHealthGoalService().GetIdByType(updateForm.HG_TYPE)
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, "error retrieving health goal")
		}
		existingUser.HG_ID = hg_id
	}

	existingUser.UpdatedAt = time.Now()

	updatedUser, err := s.userRepo.UpdateUser(existingUser)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "error updating user")
	}

	return updatedUser, nil
}

func (s *userService) DeleteUser(c echo.Context) error {
	tokenUser, err := s.tokenRepo.UserByToken(middleware.GetToken(c))
	if err != nil {
		return c.String(http.StatusUnauthorized, "Unauthorized")
	}

	existingUser, err := s.userRepo.GetUserById(tokenUser.Id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "error retrieving user")
	}
	return s.userRepo.DeleteUser(existingUser.ID)
}

func (s *userService) Logout(c echo.Context) error {
	token := middleware.GetToken(c)
	return s.userRepo.Logout(token)
}

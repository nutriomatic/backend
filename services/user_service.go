package services

import (
	"golang-template/dto"
	"golang-template/middleware"
	"golang-template/models"
	"golang-template/repository"
	"golang-template/utils"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type UserService interface {
	CreateUser(registerReq *dto.Register) error
	GetUserById(id string) (*dto.UserResponseToken, error)
	GetClassCalories(c echo.Context) (float64, int64, error)
	GetUserByEmail(email string) (*models.User, error)
	UpdateUser(updateForm *dto.RegisterForm, c echo.Context) error
	DeleteUser(c echo.Context) error
	Logout(c echo.Context) error
}

type userService struct {
	userRepo  repository.UserRepository
	tokenRepo repository.TokenRepository
	uploader  *ClientUploader
	alRepo    repository.ActivityLevelRepository
	hgRepo    repository.HealthGoalRepository
}

func NewUserService() UserService {
	return &userService{
		userRepo:  repository.NewUserRepositoryGORM(),
		tokenRepo: repository.NewTokenRepositoryGORM(),
		uploader:  NewClientUploader(),
		alRepo:    repository.NewActivityLevelRepositoryGORM(),
		hgRepo:    repository.NewHealthGoalRepositoryGORM(),
	}
}

func ParseUserForm(c echo.Context) (*dto.RegisterForm, error) {
	name := c.FormValue("name")
	email := c.FormValue("email")
	gender, err := strconv.Atoi(c.FormValue("gender"))
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "gender bad request")
	}
	telp := c.FormValue("telp")
	birthdate := c.FormValue("birthdate")
	height, err := strconv.ParseFloat(c.FormValue("height"), 64)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "height bad request")
	}
	weight, err := strconv.ParseFloat(c.FormValue("weight"), 64)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "weight bad request")
	}
	weightGoal, err := strconv.ParseFloat(c.FormValue("weight_goal"), 64)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "weight goal bad request")
	}

	alType, err := strconv.Atoi(c.FormValue("al_type"))
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "al type bad request")
	}

	hgType, err := strconv.Atoi(c.FormValue("hg_type"))
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "hg type bad request")
	}

	return &dto.RegisterForm{
		Name:       name,
		Email:      email,
		Gender:     int64(gender),
		Telp:       telp,
		Birthdate:  birthdate,
		Height:     height,
		Weight:     weight,
		WeightGoal: weightGoal,
		AL_TYPE:    int64(alType),
		HG_TYPE:    int64(hgType),
	}, nil
}

func (s *userService) CreateUser(registerReq *dto.Register) error {
	al_id, err := s.alRepo.GetActivityLevelIdByType(registerReq.AL_TYPE)
	if err != nil {
		return err
	}
	hg_id, err := s.hgRepo.GetIdByType(registerReq.HG_TYPE)
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

func (s *userService) GetUserByEmail(email string) (*models.User, error) {
	return s.userRepo.GetUserByEmail(email)
}

func (s *userService) GetUserById(id string) (*dto.UserResponseToken, error) {
	User, err := s.userRepo.GetUserById(id)
	if err != nil {
		return nil, err
	}

	al, err := s.alRepo.GetActivityLevelById(User.AL_ID)
	if err != nil {
		return nil, err
	}

	hg, err := s.hgRepo.GetById(User.HG_ID)
	if err != nil {
		return nil, err
	}

	response := &dto.UserResponseToken{
		Id:   User.ID,
		Name: User.Name,
		// Username:   User.Username,
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

func (s *userService) GetClassCalories(c echo.Context) (float64, int64, error) {
	tokenUser, err := s.tokenRepo.UserToken(middleware.GetToken(c))
	if err != nil {
		return 0, 0, echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}

	existingUser, err := s.userRepo.GetUserById(tokenUser.ID)
	if err != nil {
		return 0, 0, echo.NewHTTPError(http.StatusInternalServerError, "error retrieving user")
	}

	return existingUser.Calories, existingUser.Classification, nil
}

func (s *userService) UpdateUser(updateForm *dto.RegisterForm, c echo.Context) error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	tokenUser, err := s.tokenRepo.UserToken(middleware.GetToken(c))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}

	existingUser, err := s.userRepo.GetUserById(tokenUser.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error retrieving user")
	}

	updatedUser, err := ParseUserForm(c)
	if err != nil {
		return err
	}

	file, _ := c.FormFile("file")
	if file != nil {
		imagePath, err := s.uploader.ProcessImageUser(c)
		if err != nil {
			return err
		}
		s.uploader.DeleteImage(s.uploader.userPath, existingUser.Profpic)
		err = godotenv.Load(".env")
		if err != nil {
			return err
		}
		realImagePath := os.Getenv("IMAGE_PATH") + imagePath
		existingUser.Profpic = realImagePath
	}

	if updatedUser.Name != "" {
		existingUser.Name = updatedUser.Name
	}

	if updatedUser.Email != "" {
		existingUser.Email = updatedUser.Email
	}

	if updatedUser.Gender != 0 {
		existingUser.Gender = updatedUser.Gender
	}

	if updatedUser.Telp != "" {
		existingUser.Telp = updatedUser.Telp
	}

	if updatedUser.Birthdate != "" {
		existingUser.Birthdate = updatedUser.Birthdate
	}

	if updatedUser.Height != 0 {
		existingUser.Height = updatedUser.Height
	}

	if updatedUser.Weight != 0 {
		existingUser.Weight = updatedUser.Weight
	}

	if updatedUser.WeightGoal != 0 {
		existingUser.WeightGoal = updatedUser.WeightGoal
	}

	if updatedUser.AL_TYPE != 0 {
		existingUser.AL_ID, err = s.alRepo.GetActivityLevelIdByType(updatedUser.AL_TYPE)
		if err != nil {
			return err
		}
	}

	if updatedUser.HG_TYPE != 0 {
		existingUser.HG_ID, err = s.hgRepo.GetIdByType(updatedUser.HG_TYPE)
		if err != nil {
			return err
		}
	}

	AL, err := s.alRepo.GetActivityLevelById(existingUser.AL_ID)
	if err != nil {
		return err
	}

	// Calculate Age
	birthdate, err := time.Parse("2006-01-02", existingUser.Birthdate)
	if err != nil {
		return err
	}

	now := time.Now()
	years := now.Year() - birthdate.Year()
	if now.YearDay() < birthdate.YearDay() {
		years--
	}

	url := os.Getenv("PYTHON_API") + "/classify"
	requestData := &dto.UserRequest{
		Id:            existingUser.ID,
		Gender:        int(existingUser.Gender),
		Age:           float64(years),
		BodyWeight:    existingUser.Weight,
		BodyHeight:    existingUser.Height,
		ActivityLevel: int(AL.AL_TYPE),
	}

	responseData, err := SendRequest[dto.UserRequest, dto.UserResponse](url, *requestData)
	if err != nil {
		return err
	}

	existingUser.Calories = responseData.Calories
	existingUser.Classification = responseData.Classification
	return s.userRepo.UpdateUser(existingUser)

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

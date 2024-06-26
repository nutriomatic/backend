package dto

import "mime/multipart"

type Register struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	HG_TYPE  int64  `json:"hg_type"`
	AL_TYPE  int64  `json:"al_type"`
}

type (
	RegisterForm struct {
		Name       string                `json:"name"`
		Email      string                `json:"email"`
		Gender     int64                 `json:"gender"`
		Telp       string                `json:"telp"`
		Profpic    string                `json:"profpic"`
		Birthdate  string                `json:"birthdate"`
		Height     float64               `json:"height"`
		Weight     float64               `json:"weight"`
		WeightGoal float64               `json:"weight_goal"`
		AL_TYPE    int64                 `json:"al_type"`
		HG_TYPE    int64                 `json:"hg_type"`
		File       *multipart.FileHeader `json:"-"`
	}

	LoginForm struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	UserResponseToken struct {
		Id         string  `json:"id"`
		Name       string  `json:"name"`
		Email      string  `json:"email"`
		Role       string  `json:"role"`
		Gender     int64   `json:"gender"`
		Telp       string  `json:"telp"`
		Profpic    string  `json:"profpic"`
		Birthdate  string  `json:"birthdate"`
		Place      string  `json:"place"`
		Height     float64 `json:"height"`
		Weight     float64 `json:"weight"`
		WeightGoal float64 `json:"weight_goal"`
		HG_ID      string  `json:"hg_id"`
		HG_TYPE    int64   `json:"hg_type"`
		HG_DESC    string  `json:"hg_desc"`
		AL_ID      string  `json:"al_id"`
		AL_TYPE    int64   `json:"al_type"`
		AL_DESC    string  `json:"al_desc"`
		AL_VALUE   float64 `json:"al_value"`
	}

	UserRequest struct {
		Id            string  `json:"id"`
		Gender        int     `json:"gender"`
		Age           float64 `json:"age"`
		BodyWeight    float64 `json:"body_weight"`
		BodyHeight    float64 `json:"body_height"`
		ActivityLevel int     `json:"activity_level"`
	}

	UserResponse struct {
		Classification string  `json:"weight_status"`
		Calories       float64 `json:"calories"`
	}
)

var Register_Successful = "Registration was successful"
var FieldsRequired = "All user fields must be provided!"
var UsernameExists = "Username already exists."
var EmailExists = "Email already exists."
var PasswordShort = "Password is too short"
var ErrorRetrievingUser = "Error retrieving user"
var UserNotFound = "User not found"
var ErrorGeneratingToken = "Error generating token"
var ErrorSavingToken = "Error saving token"

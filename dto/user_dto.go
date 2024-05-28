package dto

type (
	RegisterForm struct {
		Name       string  `json:"name"`
		Username   string  `json:"username"`
		Email      string  `json:"email"`
		Password   string  `json:"password"`
		Gender     int64   `json:"gender"`
		Telp       string  `json:"telp"`
		Profpic    string  `json:"profpic"`
		Birthdate  string  `json:"birthdate"`
		Place      string  `json:"place"`
		Height     float64 `json:"height"`
		Weight     float64 `json:"weight"`
		WeightGoal float64 `json:"weightGoal"`
	}

	LoginForm struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

var Register_Successful = "Registration was successful"
var FieldsRequired = "All user fields must be provided!"
var UsernameExists = "Username already exists."
var EmailExists = "Email already exists."
var PasswordShort = "Password is too short"

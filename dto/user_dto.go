package dto

type (
	RegisterForm struct {
		Name     string `json:"name"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginForm struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

var Register_Successful = "Registration was successful"

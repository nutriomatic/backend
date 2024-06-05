package dto

type ALRegisterForm struct {
	ALType  int64   `json:"al_type"`
	ALDesc  string  `json:"al_desc"`
	ALValue float64 `json:"al_value"`
}

package dto

import "mime/multipart"

type ScannedNutritionForm struct {
	SN_PRODUCTNAME string                `json:"sn_name"`
	File           *multipart.FileHeader `json:"-"`
}

type SNResponse struct {
	Carbs        float64 `json:"carbs"`
	Energy       float64 `json:"energy"`
	Fat          float64 `json:"fat"`
	Grade        string  `json:"grade"`
	Protein      float64 `json:"protein"`
	Salt         float64 `json:"salt"`
	Sugar        float64 `json:"sugar"`
	SaturatedFat float64 `json:"saturated_fat"`
	Fiber        float64 `json:"fiber"`
}

type SNRequest struct {
	Url string `json:"url"`
}

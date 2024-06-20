package dto

import "mime/multipart"

type ScannedNutritionForm struct {
	SN_PRODUCTNAME string                `json:"sn_name"`
	File           *multipart.FileHeader `json:"-"`
}

type NutritionFacts struct {
	Carbs        float64 `json:"carbs"`
	Energy       float64 `json:"energy"`
	Fat          float64 `json:"fat"`
	Fiber        float64 `json:"fiber"`
	Protein      float64 `json:"protein"`
	SaturatedFat float64 `json:"saturated_fat"`
	Sodium       float64 `json:"sodium"`
	Sugar        float64 `json:"sugar"`
}

type SNResponse struct {
	Grade          string         `json:"grade"`
	NutritionFacts NutritionFacts `json:"nutrition_facts"`
}

type SNRequest struct {
	Url string `json:"url"`
}

package dto

import "mime/multipart"

type ProductRegisterForm struct {
	ProductName         string                `json:"product_name"`
	ProductPrice        float64               `json:"product_price"`
	ProductDesc         string                `json:"product_desc"`
	ProductIsShow       int64                 `json:"product_isshow"`
	ProductLemakTotal   float64               `json:"product_lemaktotal"`
	ProductProtein      float64               `json:"product_protein"`
	ProductKarbohidrat  float64               `json:"product_karbohidrat"`
	ProductGaram        float64               `json:"product_garam"`
	ProductGrade        string                `json:"product_grade"`
	ProductServingSize  float64               `json:"product_servingsize"`
	ProductExpShow      string                `json:"product_expshow"`
	ProductEnergi       float64               `json:"product_energi"`
	ProductGula         float64               `json:"product_gula"`
	ProductSaturatedFat float64               `json:"product_saturatedfat"`
	ProductFiber        float64               `json:"product_fiber"`
	PT_Type             int64                 `json:"pt_type"`
	File                *multipart.FileHeader `json:"-"`
}

type ProductResponse struct {
	ProductID           string  `json:"product_id"`
	ProductName         string  `json:"product_name"`
	ProductPrice        float64 `json:"product_price"`
	ProductDesc         string  `json:"product_desc"`
	ProductIsShow       int64   `json:"product_isshow"`
	ProductLemakTotal   float64 `json:"product_lemaktotal"`
	ProductProtein      float64 `json:"product_protein"`
	ProductKarbohidrat  float64 `json:"product_karbohidrat"`
	ProductGaram        float64 `json:"product_garam"`
	ProductGrade        string  `json:"product_grade"`
	ProductServingSize  float64 `json:"product_servingsize"`
	ProductExpShow      string  `json:"product_expshow"`
	ProductPicture      string  `json:"product_picture"`
	ProductEnergi       float64 `json:"product_energi"`
	ProductGula         float64 `json:"product_gula"`
	ProductSaturatedFat float64 `json:"product_saturatedfat"`
	ProductFiber        float64 `json:"product_fiber"`
	PT_Type             int64   `json:"pt_type"`
}

type ProductRequest struct {
	Id           string  `json:"id"`
	Energy       float64 `json:"energy"`
	Protein      float64 `json:"protein"`
	Fat          float64 `json:"fat"`
	Carbs        float64 `json:"carbs"`
	Sugar        float64 `json:"sugar"`
	Salt         float64 `json:"sodium"`
	SaturatedFat float64 `json:"saturated_fat"`
	Fiber        float64 `json:"fiber"`
}

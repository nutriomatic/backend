package dto

import "mime/multipart"

type ProductRegisterForm struct {
	ProductName        string                `json:"product_name"`
	ProductPrice       float64               `json:"product_price"`
	ProductDesc        string                `json:"product_desc"`
	ProductIsShow      bool                  `json:"product_isshow"`
	ProductLemakTotal  float64               `json:"product_lemaktotal"`
	ProductProtein     float64               `json:"product_protein"`
	ProductKarbohidrat float64               `json:"product_karbohidrat"`
	ProductGaram       float64               `json:"product_garam"`
	ProductGrade       string                `json:"product_grade"`
	ProductServingSize float64               `json:"product_servingsize"`
	ProductExpShow     string                `json:"product_expshow"`
	PT_Type            int64                 `json:"pt_type"`
	File               *multipart.FileHeader `json:"-"`
}

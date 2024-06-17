package dto

import "mime/multipart"

type ScannedNutritionForm struct {
	SN_PRODUCTNAME string                `json:"sn_name"`
	File           *multipart.FileHeader `json:"-"`
}

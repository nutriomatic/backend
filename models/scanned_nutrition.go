package models

import (
	"time"
)

type ScannedNutrition struct {
	SN_ID           string    `gorm:"type:primary_key" json:"sn_id"`
	SN_PRODUCTNAME  string    `json:"sn_productName"`
	SN_PRODUCTTYPE  string    `json:"sn_productType"`
	SN_INFO         string    `json:"sn_info"`
	SN_PICTURE      string    `json:"sn_picture"`
	SN_ENERGY       float64   `json:"sn_energy"`
	SN_PROTEIN      float64   `json:"sn_protein"`
	SN_FAT          float64   `json:"sn_fat"`
	SN_CARBOHYDRATE float64   `json:"sn_carbohydrate"`
	SN_SUGAR        float64   `json:"sn_sugar"`
	SN_SALT         float64   `json:"sn_salt"`
	SN_GRADE        string    `json:"sn_grade"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	USER_ID         string    `json:"user_id"`
	USER            User      `gorm:"foreignKey:USER_ID;references:ID" json:"-"`
}

package models

import (
	"time"
)

type ScannedNutrition struct {
	SN_ID          string        `gorm:"type:primary_key" json:"sn_id"`
	SN_PRODUCTNAME string        `gorm:"not null" json:"sn_productName"`
	SN_PRODUCTTYPE string        `gorm:"not null" json:"sn_productType"`
	SN_INFO        string        `gorm:"not null" json:"sn_info"`
	SN_PICTURE     string        `gorm:"not null" json:"sn_picture"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
	NI_ID          string        `json:"ni_id"`
	NI             NutritionInfo `gorm:"foreignKey:NI_ID;references:NI_ID"`
	USER_ID        string        `json:"user_id"`
	USER           User          `gorm:"foreignKey:USER_ID;references:ID"`
}

package models

import (
	"time"

	"github.com/google/uuid"
)

type ScannedNutrition struct {
	SN_ID          uuid.UUID     `gorm:"type:char(36);primary_key" json:"sn_id"`
	SN_PRODUCTNAME string        `gorm:"not null" json:"sn_productName"`
	SN_PRODUCTTYPE string        `gorm:"not null" json:"sn_productType"`
	SN_INFO        string        `gorm:"not null" json:"sn_info"`
	SN_PICTURE     string        `gorm:"not null" json:"sn_picture"`
	CreatedAt      time.Time     `json:"createdAt"`
	UpdatedAt      time.Time     `json:"updatedAt"`
	NI_ID          uuid.UUID     `json:"ni_id"`
	NI             NutritionInfo `gorm:"foreignKey:NI_ID;references:NI_ID"`
	USER_ID        uuid.UUID     `json:"user_id"`
	USER           User          `gorm:"foreignKey:USER_ID;references:ID"`
}

package models

import (
	"time"
)

type NutritionInfo struct {
	NI_ID     string    `gorm:"type:primaryey" json:"ni_id"`
	NI_TYPE   string    `gorm:"not null" json:"ni_type"`
	NI_TEXT   string    `gorm:"not null" json:"ni_text"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

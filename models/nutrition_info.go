package models

import (
	"time"

	"github.com/google/uuid"
)

type NutritionInfo struct {
	NI_ID     uuid.UUID `gorm:"type:char(36);primary_key" json:"ni_id"`
	NI_TYPE   string    `gorm:"not null" json:"ni_type"`
	NI_TEXT   string    `gorm:"not null" json:"ni_text"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

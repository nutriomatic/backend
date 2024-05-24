package models

import (
	"time"

	"github.com/google/uuid"
)

type ActivityLevel struct {
	AL_ID     uuid.UUID `gorm:"type:char(36);primary_key" json:"al_id"`
	AL_TYPE   int64     `gorm:"not null" json:"al_type"`
	AL_DESC   string    `gorm:"not null" json:"al_desc"`
	AL_VALUE  float64   `gorm:"not null" json:"al_value"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

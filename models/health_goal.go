package models

import (
	"time"

	"github.com/google/uuid"
)

type HealthGoal struct {
	HG_ID     uuid.UUID `gorm:"type:char(36);primary_key" json:"hg_id"`
	HG_TYPE   int64     `gorm:"not null" json:"hg_type"`
	HG_DESC   string    `gorm:"not null" json:"hg_desc"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

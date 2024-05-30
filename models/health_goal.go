package models

import (
	"time"
)

type HealthGoal struct {
	HG_ID     string    `gorm:"type:varchar(36);primaryKey" json:"hg_id"`
	HG_TYPE   int64     `gorm:"not null" json:"hg_type"`
	HG_DESC   string    `gorm:"not null" json:"hg_desc"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

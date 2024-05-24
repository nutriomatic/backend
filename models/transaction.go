package models

import (
	"time"

	"github.com/google/uuid"
)

type transaction struct {
	TSC_ID             uuid.UUID `gorm:"type:char(36);primary_key" json:"tsc_id"`
	TSC_PRICE          float64   `gorm:"not null" json:"tsc_price"`
	TSC_VIRTUALACCOUNT string    `gorm:"not null" json:"tsc_virtualaccount"`
	TSC_SUMPRODUCT     int       `gorm:"not null" json:"tsc_sumproduct"`
	TSC_START          time.Time `gorm:"not null" json:"tsc_start"`
	TSC_END            time.Time `gorm:"not null" json:"tsc_end"`
	CreatedAt          time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt          time.Time `gorm:"not null" json:"updatedAt"`
	STORE_ID           uuid.UUID `json:"store_id"`
	STORE              Store     `gorm:"foreignKey:STORE_ID;references:STORE_ID"`
}

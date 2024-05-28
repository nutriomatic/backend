package models

import (
	"time"
)

type Transaction struct {
	TSC_ID             string    `gorm:"type:primary_key" json:"tsc_id"`
	TSC_PRICE          float64   `gorm:"not null" json:"tsc_price"`
	TSC_VIRTUALACCOUNT string    `gorm:"not null" json:"tsc_virtualaccount"`
	TSC_SUMPRODUCT     int       `gorm:"not null" json:"tsc_sumproduct"`
	TSC_START          time.Time `gorm:"not null" json:"tsc_start"`
	TSC_END            time.Time `gorm:"not null" json:"tsc_end"`
	CreatedAt          time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt          time.Time `gorm:"not null" json:"updatedAt"`
	STORE_ID           string    `json:"store_id"`
	STORE              Store     `gorm:"foreignKey:STORE_ID;references:STORE_ID"`
	PAYMENT_ID         string    `json:"payment_id"`
	PAYMENT            Payment   `gorm:"foreignKey:PAYMENT_ID;references:PAYMENT_ID"`
}

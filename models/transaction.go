package models

import (
	"time"
)

type Transaction struct {
	TSC_ID             string    `gorm:"type:primary_key" json:"tsc_id"`
	TSC_PRICE          float64   `json:"tsc_price"`
	TSC_VIRTUALACCOUNT string    `json:"tsc_virtualaccount"`
	TSC_START          time.Time `json:"tsc_start"`
	TSC_END            time.Time `json:"tsc_end"`
	TSC_STATUS         string    `json:"tsc_status"`
	TSC_BUKTI          string    `json:"tsc_bukti"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
	PRODUCT_ID         string    `json:"product_id"`
	PRODUCT            Product   `gorm:"foreignKey:PRODUCT_ID;references:PRODUCT_ID" json:"-"`
	STORE_ID           string    `json:"store_id"`
	STORE              Store     `gorm:"foreignKey:STORE_ID;references:STORE_ID" json:"-"`
	PAYMENT_ID         string    `json:"payment_id"`
	PAYMENT            Payment   `gorm:"foreignKey:PAYMENT_ID;references:PAYMENT_ID" json:"-"`
}

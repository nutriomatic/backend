package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	PRODUCT_ID          uuid.UUID `gorm:"type:char(36);primary_key" json:"product_id"`
	PRODUCT_NAME        string    `gorm:"not null" json:"product_name"`
	PRODUCT_PRICE       float64   `gorm:"not null" json:"product_price"`
	PRODUCT_DESC        string    `gorm:"not null" json:"product_desc"`
	PRODUCT_ISSHOW      bool      `gorm:"not null" json:"product_isShow"`
	PRODUCT_LEMAKTOTAL  float64   `gorm:"not null" json:"product_lemakTotal"`
	PRODUCT_PROTEIN     float64   `gorm:"not null" json:"product_protein"`
	PRODUCT_KARBOHIDRAT float64   `gorm:"not null" json:"product_karbohidrat"`
	PRODUCT_GARAM       float64   `gorm:"not null" json:"product_garam"`
	CreatedAt           time.Time `json:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt"`
	STORE_ID            uuid.UUID `json:"store_id"`
	STORE               Store     `gorm:"foreignKey:STORE_ID;references:STORE_ID"`
}

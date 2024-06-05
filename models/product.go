package models

import (
	"time"
)

type Product struct {
	PRODUCT_ID          string        `gorm:"primaryKey" json:"product_id"`
	PRODUCT_NAME        string        `gorm:"not null" json:"product_name"`
	PRODUCT_PRICE       float64       `gorm:"not null" json:"product_price"`
	PRODUCT_DESC        string        `gorm:"not null" json:"product_desc"`
	PRODUCT_ISSHOW      bool          `gorm:"not null" json:"product_isShow"`
	PRODUCT_LEMAKTOTAL  float64       `gorm:"not null" json:"product_lemakTotal"`
	PRODUCT_PROTEIN     float64       `gorm:"not null" json:"product_protein"`
	PRODUCT_KARBOHIDRAT float64       `gorm:"not null" json:"product_karbohidrat"`
	PRODUCT_GARAM       float64       `gorm:"not null" json:"product_garam"`
	PRODUCT_GRADE       string        `gorm:"not null" json:"product_grade"`
	PRODUCT_SERVINGSIZE float64       `gorm:"not null" json:"product_servingSize"`
	CreatedAt           time.Time     `json:"createdAt"`
	UpdatedAt           time.Time     `json:"updatedAt"`
	STORE_ID            string        `json:"store_id"`
	STORE               Store         `gorm:"foreignKey:STORE_ID;references:STORE_ID" json:"-"`
	NI_ID               string        `json:"ni_id"`
	NUTRITIONINFO       NutritionInfo `gorm:"foreignKey:NI_ID;references:NI_ID" json:"-"`
}

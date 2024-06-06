package models

import (
	"time"
)

type Product struct {
	PRODUCT_ID          string        `gorm:"primaryKey" json:"product_id"`
	PRODUCT_NAME        string        `json:"product_name"`
	PRODUCT_PRICE       float64       `json:"product_price"`
	PRODUCT_DESC        string        `json:"product_desc"`
	PRODUCT_ISSHOW      bool          `json:"product_isShow"`
	PRODUCT_LEMAKTOTAL  float64       `json:"product_lemakTotal"`
	PRODUCT_PROTEIN     float64       `json:"product_protein"`
	PRODUCT_KARBOHIDRAT float64       `json:"product_karbohidrat"`
	PRODUCT_GARAM       float64       `json:"product_garam"`
	PRODUCT_SERVINGSIZE float64       `json:"product_servingSize"`
	PRODUCT_PICTURE     string        `json:"product_picture"`
	PRODUCT_GRADING     string        `json:"product_grade"`
	CreatedAt           time.Time     `json:"createdAt"`
	UpdatedAt           time.Time     `json:"updatedAt"`
	STORE_ID            string        `gorm:"type:varchar(36)" json:"store_id"`
	STORE               Store         `gorm:"foreignKey:STORE_ID;references:STORE_ID" json:"-"`
	NI_ID               string        `gorm:"type:varchar(36)" json:"ni_id"`
	NUTRITIONINFO       NutritionInfo `gorm:"foreignKey:NI_ID;references:NI_ID" json:"-"`
	PT_ID               string        `gorm:"type:varchar(36)" json:"pt_id"`
	PRODUCT_TYPE        ProductType   `gorm:"foreignKey:PT_ID;references:PT_ID" json:"-"`
}

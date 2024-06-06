package models

import (
	"time"
)

type Product struct {
	PRODUCT_ID          string      `gorm:"primaryKey" json:"product_id"`
	PRODUCT_NAME        string      `json:"product_name"`
	PRODUCT_PRICE       float64     `json:"product_price"`
	PRODUCT_DESC        string      `json:"product_desc"`
	PRODUCT_ISSHOW      bool        `gorm:"column:product_isshow" json:"product_isshow"`
	PRODUCT_LEMAKTOTAL  float64     `json:"product_lemaktotal"`
	PRODUCT_PROTEIN     float64     `json:"product_protein"`
	PRODUCT_KARBOHIDRAT float64     `gorm:"column:product_karbohidrat" json:"product_karbohidrat"`
	PRODUCT_GARAM       float64     `gorm:"column:product_garam" json:"product_garam"`
	PRODUCT_SERVINGSIZE float64     `json:"product_servingsize"`
	PRODUCT_PICTURE     string      `json:"product_picture"`
	PRODUCT_GRADING     string      `gorm:"column:product_grade" json:"product_grade"`
	PRODUCT_EXPSHOW     time.Time   `json:"product_expshow"`
	CreatedAt           time.Time   `gorm:"column:createdat" json:"created_at"`
	UpdatedAt           time.Time   `gorm:"column:updatedat" json:"updated_at"`
	STORE_ID            string      `gorm:"type:varchar(36)" json:"store_id"`
	STORE               Store       `gorm:"foreignKey:STORE_ID;references:STORE_ID" json:"-"`
	PT_ID               string      `gorm:"type:varchar(36)" json:"pt_id"`
	PRODUCT_TYPE        ProductType `gorm:"foreignKey:PT_ID;references:PT_ID" json:"-"`
}

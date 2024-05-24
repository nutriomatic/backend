package models

import "github.com/google/uuid"

type TransactionProduct struct {
	TP_ID      uuid.UUID   `gorm:"type:char(36);primary_key" json:"tp_id"`
	TSC_ID     uuid.UUID   `json:"tsc_id"`
	TSC        Transaction `gorm:"foreignKey:TSC_ID;references:TSC_ID"`
	PRODUCT_ID uuid.UUID   `json:"product_id"`
	PRODUCT    Product     `gorm:"foreignKey:PRODUCT_ID;references:ID"`
}

package models

type TransactionProduct struct {
	TP_ID      string      `gorm:"type:primary_key" json:"tp_id"`
	TSC_ID     string      `json:"tsc_id"`
	TSC        Transaction `gorm:"foreignKey:TSC_ID;references:TSC_ID"`
	PRODUCT_ID string      `json:"product_id"`
	PRODUCT    Product     `gorm:"foreignKey:PRODUCT_ID;references:ID"`
}

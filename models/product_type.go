package models

type ProductType struct {
	PT_ID   string `gorm:"primaryKey" json:"pt_id"`
	PT_Name string `json:"pt_name"`
	PT_TYPE int64  `json:"pt_type"`
}

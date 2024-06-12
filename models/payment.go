package models

import (
	"time"
)

type Payment struct {
	PAYMENT_ID     string    `gorm:"type:primary_key" json:"payment_id"`
	PAYMENT_METHOD string    `json:"payment_method"`
	PAYMENT_LINK   string    `json:"payment_link"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

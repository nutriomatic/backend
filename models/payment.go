package models

import (
	"time"
)

type Payment struct {
	PAYMENT_ID     string    `gorm:"type:primary_key" json:"payment_id"`
	PAYMENT_METHOD string    `gorm:"not null" json:"payment_method"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

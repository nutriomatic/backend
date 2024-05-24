package models

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	PAYMENT_ID     uuid.UUID `gorm:"type:char(36);primary_key" json:"payment_id"`
	PAYMENT_METHOD string    `gorm:"not null" json:"payment_method"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

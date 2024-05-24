package models

import (
	"time"

	"github.com/google/uuid"
)

type Store struct {
	STORE_ID       uuid.UUID `gorm:"type:char(36);primary_key" json:"store_id"`
	STORE_NAME     string    `gorm:"not null" json:"store_name"`
	STORE_USERNAME string    `gorm:"not null" json:"store_username"`
	STORE_CONTACT  string    `gorm:"not null" json:"store_contact"`
	CreatedAt      time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt      time.Time `gorm:"not null" json:"updatedAt"`
	USER_ID        uuid.UUID `json:"user_id"`
	USER           User      `gorm:"foreignKey:USER_ID;references:ID"`
}

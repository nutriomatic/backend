package models

import (
	"time"
)

type Store struct {
	STORE_ID       string    `gorm:"type:varchar(36);primaryKey" json:"store_id"`
	STORE_NAME     string    `gorm:"not null" json:"store_name"`
	STORE_USERNAME string    `gorm:"not null" json:"store_username"`
	STORE_ADDRESS  string    `gorm:"not null" json:"store_address"`
	STORE_CONTACT  string    `gorm:"not null" json:"store_contact"`
	CreatedAt      time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt      time.Time `gorm:"not null" json:"updatedAt"`
	USER_ID        string    `gorm:"type:varchar(36)" json:"user_id"`
	USER           User      `gorm:"foreignKey:USER_ID;references:ID"`
}

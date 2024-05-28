package models

import (
	"time"

	"gorm.io/gorm"
)

type Token struct {
	ID        string         `gorm:"primaryKey;type:varchar(36);not null" json:"id"`
	UserId    string         `gorm:"type:varchar(36);not null" json:"userId"`
	User      User           `gorm:"foreignKey:UserId;references:ID;constraint:OnDelete:CASCADE;" json:"-"`
	Token     string         `gorm:"unique;size:255;not null" json:"token"`
	ExpiresAt time.Time      `gorm:"not null" json:"expiresAt"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

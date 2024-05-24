package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	ID        uuid.UUID `gorm:"primaryKey;type:char(36);not null" json:"id"`
	UserId    uuid.UUID `gorm:"type:char(36);not null" json:"userId"`
	User      User      `gorm:"foreignKey:UserId;references:ID;constraint:OnDelete:CASCADE;" json:"-"`
	Token     string    `gorm:"unique;size:255" json:"token"`
	ExpiresAt time.Time `json:"expiresAt"`
}

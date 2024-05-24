package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID     `gorm:"type:char(36);primary_key" json:"id"`
	Username      string        `gorm:"unique;not null" json:"username"`
	Name          string        `gorm:"not null" json:"name"`
	Email         string        `gorm:"unique;not null" json:"email"`
	Password      string        `gorm:"not null" json:"password"`
	Role          string        `gorm:"not null" json:"role"`
	Gender        int64         `gorm:"not null" json:"gender"`
	Telp          string        `gorm:"not null" json:"telp"`
	Profpic       string        `gorm:"not null" json:"profpic"`
	Birthdate     string        `gorm:"not null" json:"birthdate"`
	Place         string        `gorm:"not null" json:"place"`
	Height        float64       `gorm:"not null" json:"height"`
	Weight        float64       `gorm:"not null" json:"weight"`
	WeightGoal    float64       `gorm:"not null" json:"weight_goal"`
	CreatedAt     time.Time     `json:"createdAt"`
	UpdatedAt     time.Time     `json:"updatedAt"`
	HG_ID         uuid.UUID     `json:"hg_id"`
	HEALTHGOAL    HealthGoal    `gorm:"foreignKey:HG_ID;references:HG_ID"`
	AL_ID         uuid.UUID     `json:"al_id"`
	ACTIVITYLEVEL ActivityLevel `gorm:"foreignKey:AL_ID;references:AL_ID"`
}

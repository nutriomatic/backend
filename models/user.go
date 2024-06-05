package models

import (
	"time"
)

type User struct {
	ID            string        `gorm:"type:varchar(36);primaryKey" json:"id"`
	Username      string        `gorm:"type:varchar(191)" json:"username"`
	Name          string        `gorm:"not null" json:"name"`
	Email         string        `gorm:"unique;not null" json:"email"`
	Password      string        `gorm:"not null" json:"password"`
	Role          string        `json:"role"`
	Gender        int64         `json:"gender"`
	Telp          string        `json:"telp"`
	Profpic       string        `json:"profpic"`
	Birthdate     string        `json:"birthdate"`
	Place         string        `json:"place"`
	Height        float64       `json:"height"`
	Weight        float64       `json:"weight"`
	WeightGoal    float64       `json:"weightGoal"`
	Calories      float64       `json:"calories"`
	CreatedAt     time.Time     `json:"createdAt"`
	UpdatedAt     time.Time     `json:"updatedAt"`
	HG_ID         string        `gorm:"type:varchar(36)" json:"hg_id"`
	HEALTHGOAL    HealthGoal    `gorm:"foreignKey:HG_ID;references:HG_ID"`
	AL_ID         string        `gorm:"type:varchar(36)" json:"al_id"`
	ACTIVITYLEVEL ActivityLevel `gorm:"foreignKey:AL_ID;references:AL_ID"`
}

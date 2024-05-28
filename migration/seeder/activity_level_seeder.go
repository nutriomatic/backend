package seeder

import (
	"encoding/json"
	"golang-template/models"
	"io"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeederActivityLevel(db *gorm.DB) {
	// db.AutoMigrate(&models.User{})

	file, err := os.Open("migration/json/activity_level.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	var activityLevel []models.ActivityLevel
	json.Unmarshal(byteValue, &activityLevel)

	for _, al := range activityLevel {
		al.AL_ID = uuid.New().String() // Generate new UUID
		al.CreatedAt = time.Now()      // Set current time for CreatedAt
		al.UpdatedAt = time.Now()      // Set current time for UpdatedAt

		if err := db.Create(&al).Error; err != nil {
			print(err)
		}
	}

	log.Println("Users seeded successfully")
}

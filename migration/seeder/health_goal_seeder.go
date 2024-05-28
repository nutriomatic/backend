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

func SeederHealthGoal(db *gorm.DB) {
	// db.AutoMigrate(&models.User{})

	file, err := os.Open("migration/json/health_goal.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	var healthGoal []models.HealthGoal
	json.Unmarshal(byteValue, &healthGoal)

	for _, hg := range healthGoal {
		hg.HG_ID = uuid.New().String() // Generate new UUID
		hg.CreatedAt = time.Now()      // Set current time for CreatedAt
		hg.UpdatedAt = time.Now()      // Set current time for UpdatedAt

		if err := db.Create(&hg).Error; err != nil {
			print(err)
		}
	}

	log.Println("Users seeded successfully")
}

package seeder

import (
	"encoding/json"
	"golang-template/models"
	"golang-template/utils"
	"io"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeederUser(db *gorm.DB) {
	// db.AutoMigrate(&models.User{})

	file, err := os.Open("migration/json/user.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	var users []models.User
	json.Unmarshal(byteValue, &users)

	for _, user := range users {
		user.ID = uuid.New() // Generate new UUID
		user.Password, err = utils.HashPassword(user.Password)
		if err != nil {
			log.Printf("Could not hash password for user %s: %v", user.Username, err)
		}
		user.CreatedAt = time.Now() // Set current time for CreatedAt
		user.UpdatedAt = time.Now() // Set current time for UpdatedAt

		if err := db.Create(&user).Error; err != nil {
			log.Printf("Could not seed user %s: %v", user.Username, err)
		}
	}

	log.Println("Users seeded successfully")
}

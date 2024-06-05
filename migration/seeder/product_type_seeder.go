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

func SeederProductType(db *gorm.DB) {
	// db.AutoMigrate(&models.User{})

	file, err := os.Open("migration/json/product_type.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	var productType []models.ProductType
	json.Unmarshal(byteValue, &productType)

	for _, pt := range productType {
		pt.PT_ID = uuid.New().String() // Generate new UUID
		pt.CreatedAt = time.Now()      // Set current time for CreatedAt
		pt.UpdatedAt = time.Now()      // Set current time for UpdatedAt

		if err := db.Create(&pt).Error; err != nil {
			print(err)
		}
	}

	log.Println("Users seeded successfully")
}

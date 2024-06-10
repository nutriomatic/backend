package config

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"golang-template/models"
)

func InitDB() *gorm.DB {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	logrus.Error("Error loading env file:", err)
	// }

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbName)

	newLogger := logger.New(
		logrus.New(), // Use logrus for logging
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(dbURI), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := UpdateProductIsShow(db); err != nil {
		log.Fatalf("Failed to update product_isShow: %v", err)
	}

	fmt.Println("Database connection successful")
	return db
}

func UpdateProductIsShow(db *gorm.DB) error {
	var products []models.Product
	currentDate := time.Now().Truncate(24 * time.Hour)

	// Find products where product_expShow matches the current date
	err := db.Where("product_expShow = ?", currentDate).Find(&products).Error
	if err != nil {
		return err
	}

	for _, product := range products {
		product.PRODUCT_ISSHOW = false
		if err := db.Save(&product).Error; err != nil {
			return err
		}
	}

	return nil
}

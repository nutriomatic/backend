package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"golang-template/models"
)

func InitDB() *gorm.DB {
	err1 := godotenv.Load(".env")
	if err1 != nil {
		logrus.Error("Error loading env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	// dbURI := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v TimeZone=Asia/Jakarta", dbHost, dbUser, dbPassword, dbName, dbPort)
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	// var err error
	// db, err := gorm.Open(postgres.New(postgres.Config{
	// 	DSN:                  dbURI,
	// 	PreferSimpleProtocol: true,
	// }), &gorm.Config{})

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
		panic(err)
	}

	if err := UpdateProductIsShow(db); err != nil {
		log.Fatalf("Failed to update product_isShow: %v", err)
	}
	// err = db.AutoMigrate(&models.Store{}, &models.HealthGoal{}, &models.ActivityLevel{}, &models.User{}, &models.ProductType{}, &models.Product{}) // , &models.Theater{}, &models.Screening{}, , &models.Employee{}

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

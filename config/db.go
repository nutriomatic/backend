package config

import (
	"fmt"
	"golang-template/models"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	db, err := gorm.Open(mysql.Open(dbURI), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Token{}) // , &models.Theater{}, &models.Screening{}, , &models.Employee{}

	if err != nil {
		panic(err)
	}

	fmt.Println("Database connection successful")
	return db
}

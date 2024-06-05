package main

import (
	"fmt"
	"golang-template/config"
	"golang-template/migration/seeder"
	"log"
)

func main() {
	// Initialize database connection
	db := config.InitDB()
	if db == nil {
		log.Fatal("Failed to initialize database")
	}

	// Run seeder function
	// seeder.SeederUser(db)
	// seeder.SeederActivityLevel(db)
	// seeder.SeederHealthGoal(db)
	seeder.SeederProductType(db)

	fmt.Println("Seeder executed successfully")
}

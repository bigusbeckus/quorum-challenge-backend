package main

import (
	"log"

	"github.com/bigusbeckus/quorum-challenge-backend/internal/app"
	"github.com/bigusbeckus/quorum-challenge-backend/internal/pkg/config"
	"github.com/bigusbeckus/quorum-challenge-backend/internal/pkg/database"
)

func Init() {
	log.Println("Performing initialization tasks")

	// Load config
	log.Print("Loading config file...")
	if err := config.Load(); err != nil {
		log.Fatalf("\n%v\n", err.Error())
	}
	log.Println("Done")

	log.Println("Initialization complete")
}

func SetupDB() {
	// Create gorm connection pool
	if err := database.Initialize(); err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	// Run database migrations
	log.Println("Database migrations started")
	if err := database.RunMigrations(); err != nil {
		log.Fatalf("\n%v\n", err.Error())
	}
	log.Println("Database migrations successful")

	// Run database extensions
	log.Println("Database extensions started")
	if err := database.RunExtensions(); err != nil {
		log.Fatalf("\n%v\n", err.Error())
	}
	log.Println("Database extensions successful")

	// Seed data
	log.Println("Started seeding database")
	if err := database.Seed(); err != nil {
		log.Fatalf("\n%v\n", err.Error())
	}
	log.Println("Database seeded successfully")
}

func main() {
	Init()
	SetupDB()

	app.StartServer()
}

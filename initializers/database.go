package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := os.Getenv("DB_URL")

	if dsn == "" {
		log.Fatal("DB_URL is not set in the environment")
	}

	// Correctly update the global DB variable
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	fmt.Println("✅ Connected to database:", DB.Dialector.Name())
}

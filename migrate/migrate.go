package main

import (
	"fmt"
	"log"

	"github.com/samantamayuri/Crud-go/initializers"
	"github.com/samantamayuri/Crud-go/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()

	if initializers.DB == nil {
		log.Fatal("Database connection is nil!")
	}
}

func main() {
	if initializers.DB == nil {
		log.Fatal("Database connection is nil! Cannot proceed with migration.")
	}

	err := initializers.DB.AutoMigrate(&models.Post{}, &models.Comment{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	fmt.Println("✅ Migration completed successfully!")
}

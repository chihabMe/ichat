package main

import (
	"fmt"
	"log"

	"github.com/chihabMe/ichat/server/internal/app/config"
	"github.com/chihabMe/ichat/server/internal/app/database"
	"github.com/chihabMe/ichat/server/internal/app/models"
)

func Migrate() error {
	// Load environment variables
	if err := config.InitDotenv(); err != nil {
		return fmt.Errorf("failed to load environment variables: %v", err)
	}

	// Initialize configuration
	cfg := config.InitConfig()
	fmt.Println("migrating to ",cfg.DBHost , "running on port ",cfg.Port)

	// Initialize database
	db, err := database.InitDb(cfg)
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %v", err)
	}

	// Auto-migrate models
	if err := db.AutoMigrate(&models.User{}, &models.Token{}, &models.Profile{},&models.PrivateMessage{},&models.GroupMessage{},&models.Group{}); err != nil {
		return fmt.Errorf("failed to migrate models: %v", err)
	}

	fmt.Println("Database migrated successfully")
	return nil
}

func main() {
	if err := Migrate(); err != nil {
		log.Fatal(err)
	}
}

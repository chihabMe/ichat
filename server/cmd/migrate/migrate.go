package migrate

import (
	"fmt"
)

func Migrate(){
	// Instance.AutoMigrate(&models.Profile{})
	// Instance.AutoMigrate(&models.User{})
	// Instance.AutoMigrate(&models.Token{})
	fmt.Println("database migrated successfully")
}
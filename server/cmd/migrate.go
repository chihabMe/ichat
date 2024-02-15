package cmd

import (
	"fmt"

	"github.com/chihabMe/ichat/server/internal/app/models"
)

func Migrate(){
	Instance.AutoMigrate(&models.Profile{})
	Instance.AutoMigrate(&models.User{})
	Instance.AutoMigrate(&models.Token{})
	fmt.Println("database migrated successfully")
}
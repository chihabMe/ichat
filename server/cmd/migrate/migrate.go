package main

import (
	"fmt"
	"log"

	"github.com/chihabMe/ichat/server/internal/app/config"
	"github.com/chihabMe/ichat/server/internal/app/database"
	"github.com/chihabMe/ichat/server/internal/app/models"
)

func Migrate(){
	if err :=config.InitDotenv();err!=nil{
		log.Fatal(err)
	}
	cfg := config.InitConfig()
	db,err:= database.InitDb(cfg)
	if err!=nil{
		log.Fatal(err)
		panic(err)
	}
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Token{})
	db.AutoMigrate(&models.Profile{})
	fmt.Println("database migrated successfully")
}


func main(){
Migrate()
}
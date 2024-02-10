package core

import (
	"fmt"

	"github.com/chihabMe/ichat/server/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


var Instance *gorm.DB

func ConnectDb(){
	db,err :=gorm.Open(sqlite.Open("test.db"),&gorm.Config{})
	if err!=nil{
		fmt.Println("failed to connect with the database")
		panic(err)
	}
	fmt.Println("connected to the database successfully")
	Instance=db
	Migrate()
}
func Migrate(){
	Instance.AutoMigrate(&models.User{})
}
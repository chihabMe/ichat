package core

import (
	"fmt"

	"github.com/chihabMe/ichat/server/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var Instance *gorm.DB

func ConnectDb(){
	user:=Config("DB_USER")
	password:=Config("DB_PASSWORD")
	host:=Config("DB_HOST")
	dbName:=Config("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, dbName)
	fmt.Println("connecting to ",dsn)

	db,err :=gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err!=nil{
		fmt.Println("failed to connect with the database")
		panic(err)
	}
	fmt.Println("connected to the database successfully")
	Instance=db
	Migrate()
}
func Migrate(){
	Instance.AutoMigrate(&models.Profile{})
	Instance.AutoMigrate(&models.User{})
	Instance.AutoMigrate(&models.Token{})
	fmt.Println("database migrated successfully")
}
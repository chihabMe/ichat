package database

import (
	"fmt"

	"github.com/chihabMe/ichat/server/internal/app/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var instance *gorm.DB

func InitDb(){
	user:=config.GetEnv("DB_USER")
	password:=config.GetEnv("DB_PASSWORD")
	host:=config.GetEnv("DB_HOST")
	dbName:=config.GetEnv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, dbName)
	fmt.Println("connecting to ",dsn)

	db,err :=gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err!=nil{
		fmt.Println("failed to connect with the database")
		panic(err)
	}
	instance = db
}
func GetDb() *gorm.DB {
	return instance
}
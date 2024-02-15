package database

import (
	"fmt"

	"github.com/chihabMe/ichat/server/internal/app/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var instance *gorm.DB

func InitDb(cfg *config.Config){
	dsn := cfg.DatabaseDSN()
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
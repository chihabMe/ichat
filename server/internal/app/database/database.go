package database

import (
	"fmt"

	"github.com/chihabMe/ichat/server/internal/app/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var instance *gorm.DB

func InitDb(cfg *config.Config)error{
	dsn := cfg.DatabaseDSN()
	fmt.Println("connecting to ",dsn)
	db,err :=gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err!=nil{
		return err
	}
	instance = db
	return nil
}
func GetDb() *gorm.DB {
	return instance
}
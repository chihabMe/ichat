package database

import (
	"fmt"

	"github.com/chihabMe/ichat/server/internal/app/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)




func InitDb(cfg *config.Config)(*gorm.DB,error){
	dsn := cfg.DatabaseDSN()
	fmt.Println("connecting to ",dsn)
	db,err :=gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err!=nil{
		return nil,err
	}
	return db,nil
}

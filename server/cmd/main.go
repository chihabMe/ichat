package main

import (
	"log"

	"github.com/chihabMe/ichat/server/internal/app/config"
	"github.com/chihabMe/ichat/server/internal/app/database"
	"github.com/chihabMe/ichat/server/internal/app/server"
)


func main() {
	cfg := config.InitConfig()
	if err:= database.InitDb(cfg);err!=nil{
		log.Println("failed to connect with the database")
		panic(err)
	}
	server.InitServer(cfg)
}
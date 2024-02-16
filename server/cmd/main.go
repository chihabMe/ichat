package main

import (
	"log"

	"github.com/chihabMe/ichat/server/internal/app/config"
	"github.com/chihabMe/ichat/server/internal/app/database"
	"github.com/chihabMe/ichat/server/internal/app/server"
)


func main() {
	if err := config.InitDotenv();err!=nil{
		log.Println(err)
	}
	cfg := config.InitConfig()
	db,err:= database.InitDb(cfg)
	 if err!=nil{
		log.Fatalf("failed to connect to the database: %v", err)
		panic(err)
	}
	server := server.CreateServer(cfg,db)
	server.Start()
}
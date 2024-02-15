package main

import (
	"github.com/chihabMe/ichat/server/internal/app/config"
	"github.com/chihabMe/ichat/server/internal/app/database"
	"github.com/chihabMe/ichat/server/internal/app/server"
)


func main() {
	cfg := config.InitConfig()
	database.InitDb(cfg)
	server.InitServer(cfg)
}
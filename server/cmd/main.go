package main

import (
	"github.com/chihabMe/ichat/server/internal/app/database"
	"github.com/chihabMe/ichat/server/internal/app/server"
)



func main() {
	database.InitDb()
	server.Start()
}
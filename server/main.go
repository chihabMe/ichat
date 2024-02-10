package main

import (
	"log"

	"github.com/chihabMe/ichat/server/core"
	"github.com/chihabMe/ichat/server/router"
	"github.com/gofiber/fiber/v2"
)
func setupRoutes(c *fiber.App){
	api := c.Group("/api")
	router.SetupAuthRoutes(api)
	router.SetupAccountsRoutes(api)

}

func main() {
	app := fiber.New()
	core.ConnectDb()
	setupRoutes(app)
	log.Fatal(app.Listen(":8000"))
}
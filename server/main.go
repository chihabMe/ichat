package main

import (
	"log"

	"github.com/chihabMe/ichat/server/router"
	"github.com/gofiber/fiber/v2"
)
func setupRoutes(c *fiber.App){
	api := c.Group("/api")
	router.SetupAuthRoutes(api)

}

func main() {
	app := fiber.New()
	setupRoutes(app)
	log.Fatal(app.Listen(":8000"))
}
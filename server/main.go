package main

import (
	"log"

	"github.com/chihabMe/ichat/server/core"
	"github.com/chihabMe/ichat/server/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)
func setupRoutes(app *fiber.App){
	api := app.Group("/api")
	router.SetupAuthRoutes(api)
	router.SetupAccountsRoutes(api)
}
func setupMiddleware(app *fiber.App){
app.Use(logger.New(logger.Config{

    Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
}))
}



func main() {
	app := fiber.New()
	core.ConnectDb()
	setupMiddleware(app)
	setupRoutes(app)
	log.Fatal(app.Listen(":8001"))
}
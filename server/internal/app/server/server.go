package server

import (
	"fmt"
	"log"

	"github.com/chihabMe/ichat/server/internal/app/config"
	"github.com/chihabMe/ichat/server/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("/api")
	router.SetupAuthRoutes(api)
	router.SetupAccountsRoutes(api)
}

func setupMiddleware(app *fiber.App) {
	app.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
	}))
}

func InitServer(cfg *config.Config) {
	app := fiber.New()
	setupMiddleware(app)
	setupRoutes(app)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", cfg.Port)))
}

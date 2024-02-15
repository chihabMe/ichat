package server

import (
	"fmt"
	"log"

	"github.com/chihabMe/ichat/server/internal/app/config"
	"github.com/chihabMe/ichat/server/internal/app/repositories"
	"github.com/chihabMe/ichat/server/internal/app/router"
	"github.com/chihabMe/ichat/server/internal/app/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App,db *gorm.DB) {
	api := app.Group("/api")

	//setting up the accounts app
	userRepository  :=repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	router.SetupAccountsRoutes(api,userService)
	//setting up the auth app
	router.SetupAuthRoutes(api)
}

func setupMiddleware(app *fiber.App) {
	app.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
	}))
}

func InitServer(cfg *config.Config,db *gorm.DB) {
	app := fiber.New()
	setupMiddleware(app)
	setupRoutes(app,db)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", cfg.Port)))
}

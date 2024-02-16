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

type Server struct{
	cfg *config.Config
	db *gorm.DB
	app *fiber.App
}
func CreateServer(cfg *config.Config,db *gorm.DB) *Server{
	return &Server{
		cfg: cfg,
		db: db,
		app:fiber.New() ,
	}


}

func (s *Server) setupRoutes(app *fiber.App,db *gorm.DB) {
	api := app.Group("/api")
	// creating the repositories
	userRepository  :=repositories.NewUserRepository(db)
	tokenRepository  :=repositories.NewTokenRepository(db)


	//creating the services
	userService := services.NewUserService(userRepository)
	authService := services.NewAuthService(tokenRepository)

	//creating the router
	router := router.NewRouter(authService,userService)

	router.SetupAccountsRoutes(api)
	router.SetupAuthRoutes(api)
}

func (s *Server) setupGlobalMiddleware(app *fiber.App) {
	app.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
	}))
}


func (s *Server) Start(){
	s.setupGlobalMiddleware(s.app)
	 s.setupRoutes(s.app,s.db)
	log.Fatal(s.app.Listen(fmt.Sprintf(":%s", s.cfg.Port)))
}

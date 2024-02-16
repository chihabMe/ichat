package server

import (
	"fmt"
	"log"

	"github.com/chihabMe/ichat/server/internal/app/config"
	"github.com/chihabMe/ichat/server/internal/app/errorutil"
	"github.com/chihabMe/ichat/server/internal/app/middleware"
	"github.com/chihabMe/ichat/server/internal/app/repositories"
	"github.com/chihabMe/ichat/server/internal/app/router"
	"github.com/chihabMe/ichat/server/internal/app/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

type Server struct{
	cfg *config.Config
	db *gorm.DB
	app *fiber.App
}
func errorsHandler(ctx *fiber.Ctx,err error)error{
		code := fiber.StatusInternalServerError
		var message string
		var errorData interface{}  = nil
		switch e :=err.(type){
		case *errorutil.CustomError:
			code  = e.Status
			message = e.Message
		case *errorutil.BodyValidationError:
			code = e.Status
			message  =e.Message
			errorData = e.Errors
		default:
			if fe,ok :=err.(*fiber.Error);ok{
				code  = fe.Code
				message = fe.Message
			}else{
				message = "Internal server error"
			}
		}
		if errorData==nil{
		return ctx.Status(code).JSON(fiber.Map{"status":"error","message":message})

		}
		return ctx.Status(code).JSON(fiber.Map{"status":"error","message":message,"errors":errorData})
	}
func CreateServer(cfg *config.Config,db *gorm.DB) *Server{
	return &Server{
		cfg: cfg,
		db: db,
		app:fiber.New(fiber.Config{
			ErrorHandler: errorsHandler,
		}) ,
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
	//middleware
	middleware := middleware.NewMiddleware(userService)

	//creating the router
	router := router.NewRouter(authService,userService,middleware)

	router.SetupAccountsRoutes(api)
	router.SetupAuthRoutes(api)
}

func (s *Server) setupGlobalMiddleware(app *fiber.App) {
	app.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
	}))
	app.Use(helmet.New())

}


func (s *Server) Start(){
	s.setupGlobalMiddleware(s.app)
	 s.setupRoutes(s.app,s.db)
	log.Fatal(s.app.Listen(fmt.Sprintf(":%s", s.cfg.Port)))
}

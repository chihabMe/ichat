package router

import (
	"fmt"

	"github.com/chihabMe/ichat/server/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app fiber.Router){
	router := app.Group("/auth")
	router.Post("/login",handler.Login)
	router.Get("/me",handler.Me)

	 fmt.Println(("regeared auth routes successfully"))
	
}
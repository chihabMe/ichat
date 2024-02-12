package router

import (
	"fmt"

	"github.com/chihabMe/ichat/server/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app fiber.Router){
	router := app.Group("/auth")

	router.Post("/token/obtain",handler.ObtainToken)
	router.Post("/token/logout",handler.LogoutToken)

	 fmt.Println(("regeared auth routes successfully"))
}
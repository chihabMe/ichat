package router

import (
	"fmt"

	handler "github.com/chihabMe/ichat/server/internal/app/handlers"
	"github.com/gofiber/fiber/v2"
)

func (r *Router) SetupAuthRoutes(app fiber.Router){
	authHandler := handler.NewAuthHandler(r.authService,r.userService)
	router := app.Group("/auth")
	router.Post("/token/obtain",authHandler.ObtainToken)
	router.Post("/token/logout",r.middleware.ProtectedMiddleware(),authHandler.LogoutToken)
	router.Get("/me",r.middleware.ProtectedMiddleware(),authHandler.Me)

	//
	 fmt.Println(("regeared auth routes successfully"))
}

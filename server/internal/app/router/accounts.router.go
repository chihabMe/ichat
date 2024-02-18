package router

import (
	"fmt"

	handler "github.com/chihabMe/ichat/server/internal/app/handlers"
	"github.com/gofiber/fiber/v2"
)

func (r *Router) SetupAccountsRoutes(app fiber.Router){
	accountHandler := handler.NewAccountHandler(r.userService)
	 router := app.Group("/accounts")
	 router.Get("",accountHandler.GetAllAccounts)
	 router.Post("/register",accountHandler.RegisterUser)
	 router.Post("/change-password",r.middleware.ProtectedMiddleware(),accountHandler.ChangePassword)
	 router.Get("/profile",r.middleware.ProtectedMiddleware(),accountHandler.GetAuthenticatedUserProfile)

	 fmt.Println(("regeared accounts routes successfully"))
}
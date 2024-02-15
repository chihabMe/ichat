package router

import (
	"fmt"

	"github.com/chihabMe/ichat/server/handler"
	"github.com/chihabMe/ichat/server/middleware"
	"github.com/gofiber/fiber/v2"
)


func SetupAccountsRoutes(app fiber.Router){
	 router := app.Group("/accounts")
	 router.Get("",handler.GetAllAccounts)
	 router.Post("/register",handler.Register)
	 router.Get("/profile",middleware.ProtectedMiddleware(),handler.GetAuthenticatedUserProfile)
	 router.Post("/change-password",middleware.ProtectedMiddleware(),handler.ChangePassword)

	 fmt.Println(("regeared accounts routes successfully"))
}
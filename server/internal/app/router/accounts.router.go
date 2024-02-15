package router

import (
	"fmt"

	handler "github.com/chihabMe/ichat/server/internal/app/handlers"
	"github.com/chihabMe/ichat/server/internal/app/services"
	"github.com/gofiber/fiber/v2"
)


func SetupAccountsRoutes(app fiber.Router,userService *services.UserService){
	accountHandler := handler.NewAccountHandler(userService)
	 router := app.Group("/accounts")
	//  router.Get("",handler.AccountHandler.)
	 router.Post("/register",accountHandler.RegisterUser)
	//  router.Get("/profile",middleware.ProtectedMiddleware(),handler.GetAuthenticatedUserProfile)
	//  router.Post("/change-password",middleware.ProtectedMiddleware(),handler.ChangePassword)

	 fmt.Println(("regeared accounts routes successfully"))
}
package router

import (
	"fmt"

	handler "github.com/chihabMe/ichat/server/internal/app/handlers"
	"github.com/chihabMe/ichat/server/internal/app/middleware"
	"github.com/chihabMe/ichat/server/internal/app/services"
	"github.com/gofiber/fiber/v2"
)
type Router struct {
	authService *services.AuthService
	userService *services.UserService
	middleware *middleware.Middleware
}
func   NewRouter(authService *services.AuthService,userService *services.UserService,middleware *middleware.Middleware)*Router{
	return &Router{authService: authService,userService: userService,middleware: middleware}
}
func (r *Router) SetupAuthRoutes(app fiber.Router){
	authHandler := handler.NewAuthHandler(r.authService,r.userService)
	router := app.Group("/auth")
	router.Post("/token/obtain",authHandler.ObtainToken)
	router.Post("/token/logout",r.middleware.ProtectedMiddleware(),authHandler.LogoutToken)
	router.Get("/me",r.middleware.ProtectedMiddleware(),authHandler.Me)

	//
	 fmt.Println(("regeared auth routes successfully"))
}



func (r *Router) SetupAccountsRoutes(app fiber.Router){
	accountHandler := handler.NewAccountHandler(r.userService)
	 router := app.Group("/accounts")
	 router.Get("",accountHandler.GetAllAccounts)
	 router.Post("/register",accountHandler.RegisterUser)
	 router.Post("/change-password",r.middleware.ProtectedMiddleware(),accountHandler.ChangePassword)
	 router.Get("/profile",r.middleware.ProtectedMiddleware(),accountHandler.GetAuthenticatedUserProfile)

	 fmt.Println(("regeared accounts routes successfully"))
}
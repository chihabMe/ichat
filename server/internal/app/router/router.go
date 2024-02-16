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
}
func   NewRouter(authService *services.AuthService,userService *services.UserService)*Router{
	return &Router{authService: authService,userService: userService}
}
func (r *Router) SetupAuthRoutes(app fiber.Router){
	authHandler := handler.NewAuthHandler(r.authService,r.userService)
	middlewares := middleware.NewMiddleware(r.userService)
	router := app.Group("/auth")
	router.Post("/token/obtain",authHandler.ObtainToken)
	// router.Post("/token/logout",handler.LogoutToken)
	router.Post("/me",middlewares.ProtectedMiddleware(),authHandler.Me)
	 fmt.Println(("regeared auth routes successfully"))
}



func (r *Router) SetupAccountsRoutes(app fiber.Router){
	accountHandler := handler.NewAccountHandler(r.userService)
	 router := app.Group("/accounts")
	//  router.Get("",handler.AccountHandler.)
	 router.Post("/register",accountHandler.RegisterUser)
	//  router.Get("/profile",middleware.ProtectedMiddleware(),handler.GetAuthenticatedUserProfile)
	//  router.Post("/change-password",middleware.ProtectedMiddleware(),handler.ChangePassword)

	 fmt.Println(("regeared accounts routes successfully"))
}
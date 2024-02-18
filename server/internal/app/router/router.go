package router

import (
	"github.com/chihabMe/ichat/server/internal/app/middleware"
	"github.com/chihabMe/ichat/server/internal/app/services"
)
type Router struct {
	authService *services.AuthService
	userService *services.UserService
	middleware *middleware.Middleware
}
func   NewRouter(authService *services.AuthService,userService *services.UserService,middleware *middleware.Middleware)*Router{
	return &Router{authService: authService,userService: userService,middleware: middleware}
}


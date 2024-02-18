// middleware/auth.go

package middleware

import (
	"github.com/chihabMe/ichat/server/internal/app/services"
)

type Middleware struct {
    userService *services.UserService
}

func NewMiddleware(userService *services.UserService) *Middleware {
    return &Middleware{userService: userService}
}

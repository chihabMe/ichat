// middleware/auth.go

package middleware

import (
	"log"

	"github.com/chihabMe/ichat/server/internal/app/services"
	"github.com/chihabMe/ichat/server/utils"
	"github.com/gofiber/fiber/v2"

	"github.com/golang-jwt/jwt/v5"
)

type Middleware struct {
    userService *services.UserService
}

func NewMiddleware(userService *services.UserService) *Middleware {
    return &Middleware{userService: userService}
}

func (m *Middleware) ProtectedMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
		authorizationHeader,ok := c.GetReqHeaders()["Authorization"]
        if !ok || len(authorizationHeader) == 0 {
            return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
        }
        tokenString := authorizationHeader[0]
        token, err := utils.VerifyAccessToken(tokenString)
        if err != nil {
            return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
        }
        claims := token.Claims.(jwt.MapClaims)
        userID := claims["user_id"].(string)
        user, err := m.userService.GetUserWithProfileByID(c.Context(), userID)
        if err != nil {
            log.Println(err)
            return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
        }
        c.Locals("user", user)
        return c.Next()
    }
}

package middleware

import (
	"github.com/chihabMe/ichat/server/internal/app/errorutil"
	"github.com/chihabMe/ichat/server/internal/app/models"
	"github.com/gofiber/fiber/v2"
)

func (m *Middleware) IsAdminMiddleware()fiber.Handler{
	return func(c *fiber.Ctx)error{
		user :=c.Locals("user").(*models.User)
		if user.Role!=models.AdminUserRole{
			return errorutil.ErrAuthorizedError
		}
		return c.Next()
	}
}
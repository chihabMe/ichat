package middleware

import (
	"github.com/chihabMe/ichat/server/services"
	utils "github.com/chihabMe/ichat/server/utils/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)
func ProtectedMiddleware()fiber.Handler{
	return func(c *fiber.Ctx)error{
		authorizationHeader,ok := c.GetReqHeaders()["Authorization"]
		if(!ok||len(authorizationHeader)==0){
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		tokenString := authorizationHeader[0]
		token,err :=utils.VerifyAccessToken(tokenString)
		if err!=nil{
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		claims := token.Claims.(jwt.MapClaims)
		user_id := uint(claims["user_id"].(float64))

		user,err:= services.GetUserByID(user_id)
		if(err!=nil){
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		c.Locals("user",user)
		return c.Next()
	}
}
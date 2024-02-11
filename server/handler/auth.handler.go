package handler

import "github.com/gofiber/fiber/v2"
func Me(c *fiber.Ctx)error{
	return c.JSON(fiber.Map{"message":"user"})

}
func ObtainToken(c *fiber.Ctx)error{
	type LoginInput
	return c.JSON(fiber.Map{"success":true,"token":"token data"})
}
func LogoutToken(c *fiber.Ctx)error{
	return c.JSON(fiber.Map{"success":true})
}
func VerifyToken(c *fiber.Ctx)error{
	return c.JSON(fiber.Map{"success":true})
}
package handler

import "github.com/gofiber/fiber/v2"
func Me(c *fiber.Ctx)error{
	return c.JSON(fiber.Map{"message":"user"})

}
func Login(c *fiber.Ctx)error{
	return c.JSON(fiber.Map{"message":"logged in"})
}
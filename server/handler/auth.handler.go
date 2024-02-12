package handler

import (
	"github.com/chihabMe/ichat/server/models"
	"github.com/chihabMe/ichat/server/services"
	utils "github.com/chihabMe/ichat/server/utils/jwt"
	validators "github.com/chihabMe/ichat/server/utils/validators"
	"github.com/gofiber/fiber/v2"
)
func Me(c *fiber.Ctx)error{
	return c.JSON(fiber.Map{"message":"user"})

}
func ObtainToken(c *fiber.Ctx)error{
	type LoginInput struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}
	type UserData struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	input := new(LoginInput)
	var userData UserData
	if err:=c.BodyParser(&input);err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status":"error","message":"missed fields"})
	}
	email := input.Email
	pass := input.Password
	user,_ :=new(models.User),*new(error)
	if !validators.IsEmail(email){
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status":"error","message":"invalid email"})
	}
	user,err := services.GetUserByEmail(email)
	if(err!=nil){
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status":"error","message":"unable to obtain token"})
	}
	if(user==nil){
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status":"error","message":"Invalid password or email"})
	}
	userData = UserData{
		ID: user.ID,
		Email: user.Email,
		Username: user.Username,
		Password: user.Password,
	}
	if !(validators.ComparePassword(userData.Password,pass)){
		return c.JSON(fiber.Map{"status":"error","message":"Invalid Email or Password"})
	}
	access_token,err:= utils.GenerateAccessToken(user)
	if err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status":"error","message":"unable to obtain token"})
	}
	refresh_token,err:= utils.GenerateRefreshToken(user)
	if err!=nil{
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status":"error","message":"unable to obtain token"})
		}
	services.SaveRefreshToken(user,refresh_token.Body,refresh_token.Exp)
	return c.JSON(fiber.Map{"status":"success","message":"token obtained","tokens":fiber.Map{"access_token":access_token.Body,"refresh_token":refresh_token.Body}})
}
func LogoutToken(c *fiber.Ctx)error{
	return c.JSON(fiber.Map{"success":true})
}
func VerifyToken(c *fiber.Ctx)error{
	return c.JSON(fiber.Map{"success":true})
}
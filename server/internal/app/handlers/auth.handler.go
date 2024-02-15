package handler

import (
	"fmt"

	"github.com/chihabMe/ichat/server/internal/app/models"
	"github.com/chihabMe/ichat/server/internal/app/services"
	utils "github.com/chihabMe/ichat/server/utils/jwt"
	validators "github.com/chihabMe/ichat/server/utils/validators"
	"github.com/gofiber/fiber/v2"
)
func ObtainToken(c *fiber.Ctx)error{
	input := new(schemas.LoginInput)
	var userData schemas.UserData
	if err:=c.BodyParser(&input);err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status":"error","message":"failed to parse the data"})
	}
	if err :=input.Validate();err!=nil{
		return c.Status(fiber.ErrBadRequest.Code).JSON(
			fiber.Map{
				"status":"error",
				"errors":err,
		})
	}
	email := input.Email
	pass := input.Password
	user,_ :=new(models.User),*new(error)
	user,err := services.GetUserByEmail(email)
	if(err!=nil){
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status":"error","message":"unable to obtain token"})
	}
	if(user==nil){
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status":"error","message":"Invalid password or email"})
	}
	userData = schemas.UserData{
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
	var token schemas.LogoutBody
	if err:= c.BodyParser(&token);err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"refresh_token":"required "})
	}
	if err:= token.Validate();err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status":"error","errors":err})
	}
	t := token.RefreshToken
	_,err := utils.VerifyToken(t)
	if err!=nil{
		fmt.Println(err)
		return c.JSON(fiber.Map{"status":"success","message":"logged out"})
	}
	services.DeleteRefreshTokenIfExisted(t)
	return c.JSON(fiber.Map{"status":"success","message":"logged out"})
}
func VerifyToken(c *fiber.Ctx)error{
	return c.JSON(fiber.Map{"success":true})
}
func Me(c *fiber.Ctx)error{
	user := c.Locals("user").(models.User)
	return c.JSON(fiber.Map{"status":"success","user":user})
}
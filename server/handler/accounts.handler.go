package handler

import (
	"github.com/chihabMe/ichat/server/core"
	"github.com/chihabMe/ichat/server/models"
	utils "github.com/chihabMe/ichat/server/utils/validators"
	"github.com/gofiber/fiber/v2"
)


type RegisterInput struct {
		Username string `json:"username"`
		Email string `json:"email"`
		Password string `json:"password"`
		Password2 string `json:"password2"`
	}
func Register(c *fiber.Ctx)error{
	db := core.Instance
	var user models.User
	    var userInput RegisterInput
	  if err := c.BodyParser(&userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Missed fields",})
    }
	if userInput.Password!=userInput.Password2{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":"Passwords don't match",
		})
	}
	hashedPassword , err :=utils.HashPassword(userInput.Password)
	if err!=nil{
		return err
	}
	user.Username=userInput.Username
	user.Email=userInput.Email
	user.Password=hashedPassword
	if err :=db.Save(&user).Error;err!=nil{
		return err
	}
	return c.JSON(fiber.Map{
		"message":"User registered successfully",
		"user":user,
	})


}
func UpdateProfile(c *fiber.Ctx)error{
	return c.JSON(fiber.Map{"success":true,"message":"your profile has been updated"})

}
func ChangePassword(c *fiber.Ctx)error{
	return c.JSON(fiber.Map{"success":true,"message":"your password has been changed"})
}
func DeleteAccount(c *fiber.Ctx)error{

	return c.JSON(fiber.Map{"success":true,"message":"your account has been deleted"})

}
func GetAllAccounts(c *fiber.Ctx)error{
	db := core.Instance
	users := db.Model(&models.User{})
	return c.JSON(fiber.Map{"users":users,})
}
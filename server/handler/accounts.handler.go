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
	var users []models.User
	
	if err :=db.Model(&models.User{}).Preload("Tokens").Find(&users).Error; err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status":"error","message":"server error"})
	}

	return c.JSON(fiber.Map{"users":users,})
}


type MeData struct {
	Username string `json"username"`
	Email string `json"email"`
	Verified bool `json:"verified"`
	Active bool `json:"active"`
}
func Me(c *fiber.Ctx)error{
	user := c.Locals("user").(*models.User)
	 meData :=MeData{
		Username: user.Username,
		Email: user.Email,
		Verified: user.Verified,
		Active: user.Active,
	 }
	return c.JSON(meData)
}
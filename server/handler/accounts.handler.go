package handler

import (
	"fmt"

	"github.com/chihabMe/ichat/server/core"
	"github.com/chihabMe/ichat/server/models"
	"github.com/chihabMe/ichat/server/services"
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"unable to register",})
	}
	user.Username=userInput.Username
	user.Email=userInput.Email
	user.Password=hashedPassword
	if err :=services.CreateUser(user);err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"unable to register",})
	}
	 profile := models.Profile{
		User: user,
	}

	if err :=services.CreateProfile(profile);err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"unable to register",})
	}
	

	return c.JSON(fiber.Map{
		"message":"User registered successfully",
		"profile":profile,
	})


}
func UpdateProfile(c *fiber.Ctx)error{
	return c.JSON(fiber.Map{"success":true,"message":"your profile has been updated"})
}
type ChangePasswordData struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
	NewPassword2 string `json:"new_password2"`
}
func ChangePassword(c *fiber.Ctx)error{
	var changePasswordData ChangePasswordData
	if err:=c.BodyParser(&changePasswordData);err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status":"error","message":"Missed fields "})
	}
	if changePasswordData.NewPassword!=changePasswordData.NewPassword2{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status":"error","message":"Passwords don't match"})
	}
	user := c.Locals("user").(*models.User)
	isSamePassword := utils.ComparePassword(user.Password,changePasswordData.OldPassword)
	if !isSamePassword{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status":"error","message":"Invalid password"})
	}
	newPasswordHash,err:=utils.HashPassword(changePasswordData.NewPassword)
	if err!=nil{
	fmt.Println(err)
	return c.SendStatus(fiber.StatusInternalServerError)
	}

	if err:=services.UpdateUserPassword(newPasswordHash,user);err!=nil{
	fmt.Println(err)
	return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"success":true,"message":"your password has been changed"})
}
func DeleteAccount(c *fiber.Ctx)error{

	return c.JSON(fiber.Map{"success":true,"message":"your account has been deleted"})

}
func GetAllAccounts(c *fiber.Ctx)error{
	db := core.Instance
	var profiles []models.Profile
	
	if err :=db.Model(&models.Profile{}).Find(&profiles).Error; err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status":"error","message":"server error"})
	}

	return c.JSON(fiber.Map{"profiles":profiles,})
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
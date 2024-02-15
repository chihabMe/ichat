package handler

import (
	"log"
	"strings"

	"github.com/chihabMe/ichat/server/internal/app/database"
	"github.com/chihabMe/ichat/server/internal/app/models"
	"github.com/chihabMe/ichat/server/internal/app/services"
	"github.com/chihabMe/ichat/server/schemas"
	utils "github.com/chihabMe/ichat/server/utils/validators"
	"github.com/gofiber/fiber/v2"
)


func Register(c *fiber.Ctx)error{
	var user models.User
	    var userInput schemas.RegisterInput
	  if err := c.BodyParser(&userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Failed to parse data",})
    }
	err := userInput.Validate()
	if(err!=nil){
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status":"error","errors":err})
	}
	hashedPassword , err :=utils.HashPassword(userInput.Password)
	if err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"unable to register",})
	}
	user.Username=userInput.Username
	user.Email=userInput.Email
	user.Password=hashedPassword
	if err :=services.CreateUser(&user);err!=nil{
	if strings.Contains(err.Error(), "Duplicate entry") {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "This email is already being used"})
    }

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "unable to register"})
    } 
        
	 profile := models.Profile{
		UserId: user.ID,
	}

	if err :=services.CreateProfile(&profile);err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"unable to register",})
	}
	

	return c.JSON(fiber.Map{
		"message":"User registered successfully",
		"profile":profile,
	})


}

func ChangePassword(c *fiber.Ctx)error{
	var changePasswordData schemas.ChangePasswordData
	if err:=c.BodyParser(&changePasswordData);err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status":"error","message":"Failed to parse data"})
	}
	if err := changePasswordData.Validate();err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status":"error","errors":err})
	}
	
	user := c.Locals("user").(*models.User)
	isSamePassword := utils.ComparePassword(user.Password,changePasswordData.OldPassword)
	if !isSamePassword{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status":"error","message":"Invalid password"})
	}
	newPasswordHash,err:=utils.HashPassword(changePasswordData.NewPassword)
	if err!=nil{
	log.Println(err)
	return c.SendStatus(fiber.StatusInternalServerError)
	}

	if err:=services.UpdateUserPassword(newPasswordHash,user);err!=nil{
		log.Println(err)
	return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"success":true,"message":"your password has been changed"})
}
func DeleteAccount(c *fiber.Ctx)error{

	return c.JSON(fiber.Map{"success":true,"message":"your account has been deleted"})

}
func GetAllAccounts(c *fiber.Ctx)error{
	db := database.GetDb()
	var users []models.User
	
	if err :=db.Model(&models.Profile{}).Find(&users).Error; err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status":"error","message":"server error"})
	}

	return c.JSON(fiber.Map{"users":users,})
}


func GetAuthenticatedUserProfile(c *fiber.Ctx)error{
	profile := c.Locals("user").(models.User)
	return c.JSON(fiber.Map{"status":"success","data":profile})
}

func UpdateProfile(c *fiber.Ctx)error{
	user := c.Locals("user").(models.User)
	var updateProfileData  schemas.UpdateProfileData
	if err:=c.BodyParser(&updateProfileData);err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status":"error","message":"failed to parse data"})
	}
	if err:= updateProfileData.Validate();err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status":"error","message":"invalid fields","errors":err})
	}
	user.Username=updateProfileData.Username
	user.Profile.PhoneNumber=updateProfileData.PhoneNumber
	if err:=services.UpdateUser(&user);err!=nil{
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status":"error","message":"server error"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status":"success","message":"updated","data":user})
}
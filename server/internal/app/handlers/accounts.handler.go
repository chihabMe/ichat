package handler

import (
	"log"
	"strings"

	"github.com/chihabMe/ichat/server/internal/app/dto"
	"github.com/chihabMe/ichat/server/internal/app/errorutil"
	"github.com/chihabMe/ichat/server/internal/app/models"
	"github.com/chihabMe/ichat/server/internal/app/services"
	"github.com/chihabMe/ichat/server/utils"
	"github.com/gofiber/fiber/v2"
)

type AccountHandler struct {
	userService *services.UserService
}
func NewAccountHandler(userService *services.UserService)*AccountHandler{
	return &AccountHandler{userService: userService}
}

func (h *AccountHandler) RegisterUser(c *fiber.Ctx)error{
	var user models.User
	var body dto.RegisterUserRequestDTO
	ctx :=c.Context()
	  if err := c.BodyParser(&body); err != nil {
		return errorutil.ErrFailedToParseData
    }
	err := body.Validate()
	if(err!=nil){
		return errorutil.NewValidationError(err,"Invalid fields ")
	}
	hashedPassword , err :=utils.HashPassword(body.Password)
	if err!=nil{
		return errorutil.ErrInternalServerError
	}
	user.Username=body.Username
	user.Email=body.Email
	user.Password=hashedPassword
	if err := h.userService.CreateUser(ctx,&user);err!=nil{
	if strings.Contains(err.Error(), "Duplicate entry") {
		errors := map[string]string {"email":"this email is already being used"}
		return errorutil.NewValidationError(errors)
    }

	return errorutil.ErrInternalServerError
    } 
        
	response := dto.RegisterUserResponseDTO{
		BaseResponseDTO: dto.BaseResponseDTO{
			Message: "User registered  successfully",
			Status: dto.StatusSuccess,
		},
		Data: dto.RegisterUserResponseDataDTO{
			UserId: user.ID.String(),
			UserEmail: user.Email,
			UserUsername: user.Username,
		},

	}
	return c.JSON(response)
}



func (h *AccountHandler) ChangePassword(c *fiber.Ctx)error{
	var body dto.ChangePasswordRequestDTO
	if err:=c.BodyParser(&body);err!=nil{
		return errorutil.ErrFailedToParseData
	}
	if err := body.Validate();err!=nil{
		return errorutil.NewValidationError(err)
	}
	
	user := c.Locals("user").(*models.User)
	isSamePassword := utils.ComparePassword(user.Password,body.OldPassword)
	if !isSamePassword{
		errors := map[string]string {"old_password":"invalid password"}
		return errorutil.NewValidationError(errors)
	}
	newPasswordHash,err:=utils.HashPassword(body.NewPassword)
	if err!=nil{
		log.Println(err)
		return errorutil.ErrInternalServerError
	}


	ctx :=c.Context()
	if err := h.userService.UpdateUserPassword(ctx,user,newPasswordHash);err!=nil{
		log.Println(err)
		return errorutil.ErrInternalServerError
	}
	response := dto.ChangePasswordResponseDTO{
		BaseResponseDTO: dto.BaseResponseDTO{
			Message: "you password has been changed",
			Status: dto.StatusSuccess,
		},
	}
	return c.JSON(response)
}
func (h *AccountHandler) GetAllAccounts(c *fiber.Ctx)error{
	var users []models.User
	
	ctx := c.Context()
	if err:= h.userService.GetAllUsers(ctx,&users);err!=nil{
		return errorutil.ErrInternalServerError
	}
	response := dto.GetAllAccountsRespondDTO{
		BaseResponseDTO: dto.BaseResponseDTO{
			Message: "all users",
			Status: dto.StatusSuccess,
			Data: users,
		},
	}
	return c.JSON(response)
}


// func GetAuthenticatedUserProfile(c *fiber.Ctx)error{
// 	profile := c.Locals("user").(models.User)
// 	return c.JSON(fiber.Map{"status":"success","data":profile})
// }

// func UpdateProfile(c *fiber.Ctx)error{
// 	user := c.Locals("user").(models.User)
// 	var updateProfileData  schemas.UpdateProfileData
// 	if err:=c.BodyParser(&updateProfileData);err!=nil{
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status":"error","message":"failed to parse data"})
// 	}
// 	if err:= updateProfileData.Validate();err!=nil{
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status":"error","message":"invalid fields","errors":err})
// 	}
// 	user.Username=updateProfileData.Username
// 	user.Profile.PhoneNumber=updateProfileData.PhoneNumber
// 	if err:=services.UpdateUser(&user);err!=nil{
// 		log.Println(err)
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status":"error","message":"server error"})
// 	}
// 	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status":"success","message":"updated","data":user})
// }
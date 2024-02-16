package handler

import (
	"github.com/chihabMe/ichat/server/internal/app/dto"
	"github.com/chihabMe/ichat/server/internal/app/models"
	"github.com/chihabMe/ichat/server/internal/app/services"
	"github.com/chihabMe/ichat/server/utils"
	"github.com/gofiber/fiber/v2"
)
type AuthHandler struct {
	authService *services.AuthService
	userService *services.UserService
}

func NewAuthHandler(authService *services.AuthService,userService *services.UserService)*AuthHandler{
	return &AuthHandler{authService:authService,userService: userService}
}



func (h *AuthHandler) ObtainToken(c *fiber.Ctx)error{
	var body dto.ObtainTokenRequestDTO
	if err:=c.BodyParser(&body);err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status":"error","message":"failed to parse the data"})
	}
	if err :=body.Validate();err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(
			dto.ObtainTokenResponseDTO{
				BaseResponseDTO: dto.BaseResponseDTO{
					Message: "Invalid fields",
					Status: dto.StatusError,
					Errors: err,
					Data: nil,
				},
			},
		)
	}
	ctx := c.Context()
	user,err := h.userService.GetUserByEmail(ctx,body.Email)
	if err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status":"error","message":"unable to obtain token"})
	}
	if(err!=nil){
		return c.Status(fiber.StatusInternalServerError).JSON(
			dto.ObtainTokenResponseDTO{
				BaseResponseDTO: dto.BaseResponseDTO{
					Message: "internal error",
					Status: dto.StatusError,
					Errors: err,
					Data: nil,
				},
			},
		)

	}
	if !(utils.ComparePassword(user.Password,body.Password)){
		return c.Status(fiber.StatusInternalServerError).JSON(
			dto.ObtainTokenResponseDTO{
				BaseResponseDTO: dto.BaseResponseDTO{
					Message: "invalid password or email ",
					Status: dto.StatusError,
				},
			},
		)
	}
	accessToken,err:= utils.GenerateAccessToken(user)
	if err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status":"error","message":"unable to obtain token"})
	}
	refreshToken,err:= utils.GenerateRefreshToken(user)
	if err!=nil{
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status":"error","message":"unable to obtain token"})
		}
	 token := models.Token{
		Exp: refreshToken.Exp,
		Token: refreshToken.Body,
		UserID: user.ID,
	}
	h.authService.SaveRefreshToken(ctx,&token)
	return c.JSON(
		dto.ObtainTokenResponseDTO{
			BaseResponseDTO: dto.BaseResponseDTO{
				Message: "token obtained",
				Status: dto.StatusSuccess,
			},
			Data:dto.ObtainTokenResponseDataDTO{
				AccessToken: accessToken.Body,
				RefreshToken: refreshToken.Body,
			} ,
		},
	)
}

// func LogoutToken(c *fiber.Ctx)error{
// 	var token schemas.LogoutBody
// 	if err:= c.BodyParser(&token);err!=nil{
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"refresh_token":"required "})
// 	}
// 	if err:= token.Validate();err!=nil{
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status":"error","errors":err})
// 	}
// 	t := token.RefreshToken
// 	_,err := utils.VerifyToken(t)
// 	if err!=nil{
// 		fmt.Println(err)
// 		return c.JSON(fiber.Map{"status":"success","message":"logged out"})
// 	}
// 	services.DeleteRefreshTokenIfExisted(t)
// 	return c.JSON(fiber.Map{"status":"success","message":"logged out"})
// }
// func VerifyToken(c *fiber.Ctx)error{
// 	return c.JSON(fiber.Map{"success":true})
// }

func (h *AuthHandler) Me(c *fiber.Ctx)error{
	user := c.Locals("user").(models.User)
	return c.JSON(fiber.Map{"status":"success","user":user})
}
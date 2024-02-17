package handler

import (
	"fmt"

	"github.com/chihabMe/ichat/server/internal/app/dto"
	"github.com/chihabMe/ichat/server/internal/app/errorutil"
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
		return errorutil.ErrFailedToParseData
	}
	if err :=body.Validate();err!=nil{
		return errorutil.NewValidationError(err)
	}
	ctx := c.Context()
	user,err := h.userService.GetUserByEmail(ctx,body.Email)
	if err!=nil{
		errors :=map[string]string{
			"password":"Invalid password",
			"email":"Invalid invalid email",
		}
		return errorutil.NewValidationError(errors,"invalid password or email")
	}
	if !(utils.ComparePassword(user.Password,body.Password)){
		errors :=map[string]string{
			"password":"Invalid password",
			"email":"Invalid invalid email",
		}
		return errorutil.NewValidationError(errors,"invalid password or email")
	}
	accessToken,err:= utils.GenerateAccessToken(user)
	if err!=nil{
		return errorutil.ErrInternalServerError
	}
	refreshToken,err:= utils.GenerateRefreshToken(user)
	if err!=nil{
		return errorutil.ErrInternalServerError
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

func (h *AuthHandler) LogoutToken(c *fiber.Ctx)error{
	var body dto.LogoutTokenRequestDTO

	if err:= c.BodyParser(&body);err!=nil{
		return errorutil.ErrFailedToParseData
	}
	if err:= body.Validate();err!=nil{
		return errorutil.NewValidationError(err)
	}
	t := body.RefreshToken
	_,err := utils.VerifyToken(t)
	response := dto.LogoutTokenResponseDTO{
		BaseResponseDTO: dto.BaseResponseDTO{
			Message: "logged out",
			Status: dto.StatusSuccess,
		},
	}
	if err!=nil{
		fmt.Println(err)
		return c.JSON(response)
	}
	ctx := c.Context()
	h.authService.DeleteRefreshTokenIfExisted(ctx,t)
	return c.JSON(response)
}

func (h *AuthHandler) Me(c *fiber.Ctx)error{
	user := c.Locals("user").(*models.User)
	return c.JSON(fiber.Map{"status":"success","user":user})
}
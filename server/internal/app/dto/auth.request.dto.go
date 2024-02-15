package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type ObtainTokenRequestDTO struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}


func (l ObtainTokenRequestDTO) Validate()error{
	return  validation.ValidateStruct(&l,
		validation.Field(&l.Email,validation.Required,is.Email),
		validation.Field(&l.Password,validation.Required),
	)
}


type LogoutTokenRequestDTO struct {
	RefreshToken string `json:"refresh_token"`
}

func (l LogoutTokenRequestDTO) Validate()error{
	return validation.ValidateStruct(&l,
		validation.Field(&l.RefreshToken,validation.Required),
	)
}
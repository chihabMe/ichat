package schemas

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
)

type LoginInput struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}


func (l LoginInput) Validate()error{
	return  validation.ValidateStruct(&l,
		validation.Field(&l.Email,validation.Required,is.Email),
		validation.Field(&l.Password,validation.Required),
	)
}
type UserData struct {
		ID       uuid.UUID   `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}


type LogoutBody struct {
	RefreshToken string `json:"refresh_token"`
}

func (l LogoutBody) Validate()error{
	return validation.ValidateStruct(&l,
		validation.Field(&l.RefreshToken,validation.Required),
	)
}
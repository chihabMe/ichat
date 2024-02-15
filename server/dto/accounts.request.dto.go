package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type RegisterUserRequestDTO struct {
		Username string `json:"username"`
		Email string `json:"email"`
		Password string `json:"password"`
		Password2 string `json:"password2"`
	}
func (r RegisterUserRequestDTO) Validate()error{
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email,validation.Required,is.Email),
		validation.Field(&r.Username,validation.Required,validation.Length(6,30)),
		validation.Field(&r.Password,validation.Required,validation.Length(6,30)),
		validation.Field(&r.Password2,validation.Required,validation.In(r.Password).Error("Passwords don't match")),
	)
}
type ChangePasswordData struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
	NewPassword2 string `json:"new_password2"`
}
func (c ChangePasswordData) Validate()error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.OldPassword,validation.Required),
		validation.Field(&c.NewPassword,validation.Required),
		validation.Field(&c.NewPassword2,validation.Required,validation.In(&c.NewPassword).Error("Passwords don't match")),
	)
}

type MeData struct {
	Username string `json"username"`
	Email string `json"email"`
	Verified bool `json:"verified"`
	Active bool `json:"active"`
}

type UpdateProfileData struct {
	Username string `json"username"`
	PhoneNumber string `json:"phone_number"`
}

func (u UpdateProfileData) Validate() error{
	return validation.ValidateStruct(&u,
		validation.Field(&u.Username,validation.Length(6,30)),
		validation.Field(&u.PhoneNumber,validation.Length(6,30)),
	)
}
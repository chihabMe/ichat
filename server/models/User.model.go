package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Email string 
	Password string `json:"-"`
	Verified bool
	Active bool
	Tokens []Token
}
type Profile struct{
	gorm.Model
	User User `gorm:"embedded"`
	PhoneNumber string `json:"phone_number"`
}





package models

import (
	"gorm.io/gorm"
)
type Role int 
const (
	AdminUser Role =iota
	NormalUser 
)
type User struct {
	Base
	Username string
	Email string 
	Password string `json:"-"`
	Verified bool
	Active bool
	Profile Profile
	Tokens []Token
	Role Role

}

func (u *User) AfterSave(tx *gorm.DB)(err error){
	
	return 
}



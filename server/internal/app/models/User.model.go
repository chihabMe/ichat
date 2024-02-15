package models

import (
	"gorm.io/gorm"
)

type Role int

const (
	NormalUserRole Role = iota
	AdminUserRole
)

var roleStrings = [...]string{
	"normal",
	"admin",
}

func (r Role) String() string {
	if r < NormalUserRole || r > AdminUserRole {
		return "unknown"
	}
	return roleStrings[r]
}

type User struct {
	Base
	Username string
	Email    string `gorm:"unique"`
	Password string `json:"-"`
	Verified bool
	Active   bool
	Profile  Profile
	Tokens   []Token
	Role     Role
}

func (u *User) AfterSave(tx *gorm.DB) (err error) {
	return
}

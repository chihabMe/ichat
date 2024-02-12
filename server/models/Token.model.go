package models

import "gorm.io/gorm"



type Token struct {
    Token string `json:"token" gorm:"index"`
	UserID uint
	Exp int64 `json:"exp"`
    gorm.Model
}

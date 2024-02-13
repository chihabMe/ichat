package models

import "github.com/google/uuid"



type Token struct {
	Base
    Token string `json:"token" gorm:"index"`
	UserID uuid.UUID
	Exp int64 `json:"exp"`
}

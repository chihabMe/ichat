package models

import "github.com/google/uuid"

type Profile struct{
	Base
	UserId uuid.UUID `json:"user_id"`
	PhoneNumber string `json:"phone_number"`
}


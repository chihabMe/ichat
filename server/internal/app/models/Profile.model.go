package models

import "github.com/google/uuid"

type Profile struct{
	Base
	UserId uuid.UUID 
	PhoneNumber string `json:"phone_number"`
}


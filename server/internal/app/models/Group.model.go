package models

import "github.com/google/uuid"

type Group struct {
	Base
	Name      string `json:"name"`
	CreatorID uuid.UUID 
	Members  []*User `gorm:"many2many:group_users;"`
	Messages []GroupMessage `json:"messages"`
}
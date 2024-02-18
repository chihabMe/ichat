package models

import "github.com/google/uuid"

type PrivateMessage struct {
	Body string `json:"body"`
	SenderId uuid.UUID `json:"sender_id"`
	ReceiverId uuid.UUID `json:"receiver_id"`
}
type GroupMessage struct {
	Body string `json:"body"`
	SenderId uuid.UUID `json:"sender_id"`
	GroupId uuid.UUID 
}
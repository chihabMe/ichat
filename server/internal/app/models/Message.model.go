package models

import "github.com/google/uuid"

type BaseMessage  struct{
	Base
	Body string `json:"body"`
	Sender User `json:"sender"`
}
type PrivateMessage struct {
	BaseMessage
	Receiver User `json:"receiver"`
}
type GroupMessage struct {
	BaseMessage
	GroupId uuid.UUID 
}
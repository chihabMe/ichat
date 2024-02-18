package models

type Message  struct{
	Body string `json:"body"`
	Sender User `json:"sender"`
	Receiver User `json:"receiver"`
}
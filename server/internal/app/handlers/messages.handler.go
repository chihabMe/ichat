package handler

import "github.com/chihabMe/ichat/server/internal/app/services"

type MessagesHandler struct {
	MessageService services.MessageService
}
func NewMessagesHandler(messageService services.MessageService)*MessagesHandler{
	return &MessagesHandler{
		MessageService:messageService ,
	}
}


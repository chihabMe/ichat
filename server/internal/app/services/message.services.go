package services

import (
	"context"

	"github.com/chihabMe/ichat/server/internal/app/models"
	"github.com/chihabMe/ichat/server/internal/app/repositories"
)


type MessageService struct {
	messageRepository repositories.MessageRepository
}

func NewMessageService(messageRepository repositories.MessageRepository)*MessageService{
	return &MessageService{messageRepository: messageRepository}
}

func (s *MessageService) GelAllGroupMessages(ctx context.Context,groupID string)(*[]models.GroupMessage,error){
	return s.messageRepository.GetAllByGroupID(ctx,groupID);
}
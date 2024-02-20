package repositories

import (
	"context"

	"github.com/chihabMe/ichat/server/internal/app/models"
	"gorm.io/gorm"
)
type MessageRepository interface{
	GetAllByGroupID(ctx context.Context,groupId string)(*[]models.GroupMessage,error)
}

type MessageRepositoryImpl struct {
	db *gorm.DB
}
func NewMessageRepository(db *gorm.DB)MessageRepository{
	return &MessageRepositoryImpl{db: db}
}
func (m *MessageRepositoryImpl) GetAllByGroupID(ctx context.Context,groupId string)(*[]models.GroupMessage,error){
	var messages  []models.GroupMessage
	if err:= m.db.WithContext(ctx).Model(&models.GroupMessage{}).Where("group_id = ?",groupId).Error;err!=nil{
		return nil,err
	}
	return &messages,nil

}

package repository

import (
	"app/internal/model"

	"gorm.io/gorm"
)

type IChatMessageRepository interface {
	CreateChatMessage(chatMessage *model.ChatMessage) error
	GetAllChatMessageByChatRoomId(chatRoomId uint) ([]model.ChatMessage, error)
}

type chatMessageRepository struct {
	db *gorm.DB
}

func NewChatMessageRepository(db *gorm.DB) IChatMessageRepository {
	return &chatMessageRepository{db}
}

func (crr *chatMessageRepository) CreateChatMessage(chatMessage *model.ChatMessage) error {
	if err := crr.db.Create(&chatMessage).Error; err != nil {
		return err
	}
	return nil
}

func (crr *chatMessageRepository) GetAllChatMessageByChatRoomId(chatRoomId uint) ([]model.ChatMessage, error) {
	var chatMessages []model.ChatMessage
	if result := crr.db.Where("chat_room_id=?", chatRoomId).Find(&chatMessages); result.Error != nil {
		return nil, result.Error
	}
	return chatMessages, nil
}

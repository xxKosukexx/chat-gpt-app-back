package repository

import (
	"app/internal/model"

	"gorm.io/gorm"
)

type IChatMessageRepository interface {
	CreateChatMessage(chatMessage *model.ChatMessage) error
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

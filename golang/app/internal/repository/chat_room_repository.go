package repository

import (
	"app/internal/model"

	"gorm.io/gorm"
)

type IChatRoomRepository interface {
	CreateChatRoom(chatRoom *model.ChatRoom) error
}

type chatRoomRepository struct {
	db *gorm.DB
}

func NewChatRoomRepository(db *gorm.DB) IChatRoomRepository {
	return &chatRoomRepository{db}
}

func (crr *chatRoomRepository) CreateChatRoom(chatRoom *model.ChatRoom) error {
	if err := crr.db.Create(&chatRoom).Error; err != nil {
		return err
	}
	return nil
}

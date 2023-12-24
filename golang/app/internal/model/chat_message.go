package model

import (
	"gorm.io/gorm"
	"time"
)

type ChatMessage struct {
	gorm.Model
	ID         uint      `json:"id" gorm:"primaryKey"`
	Question   string    `json:"question" gorm:"not null"`
	Answer     string    `json:"answer" gorm:"not null"`
	ChatRoomId uint      `json:"chat_room_id"`
	ChatRoom   ChatRoom  `json:"chat_room" gorm:"foreignKey:ChatRoomId; constraint:OnDelete:CASCADE"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ChatMessageResponse struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Question string `json:"question" gorm:"not null"`
	Answer   string `json:"answer" gorm:"not null"`
}

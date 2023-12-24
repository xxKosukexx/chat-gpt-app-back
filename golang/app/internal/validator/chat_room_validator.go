package validator

import (
	"app/internal/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IChatRoomValidator interface {
	ChatRoomValidate(chatRoom model.ChatRoom) error
}

type chatRoomValidator struct{}

func NewChatRoomValidator() IChatRoomValidator {
	return &chatRoomValidator{}
}

func (uv *chatRoomValidator) ChatRoomValidate(chatRoom model.ChatRoom) error {
	return validation.ValidateStruct(&chatRoom,
		validation.Field(
			&chatRoom.Title,
			validation.Required.Error("title is required"),
		),
	)
}

package validator

import (
	"app/internal/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IChatMessageValidator interface {
	ChatMessageValidate(chatMessage model.ChatMessage) error
}

type chatMessageValidator struct{}

func NewChatMessageValidator() IChatMessageValidator {
	return &chatMessageValidator{}
}

func (uv *chatMessageValidator) ChatMessageValidate(chatMessage model.ChatMessage) error {
	return validation.ValidateStruct(&chatMessage,
		validation.Field(
			&chatMessage.Question,
			validation.Required.Error("question is required"),
		),
	)
}

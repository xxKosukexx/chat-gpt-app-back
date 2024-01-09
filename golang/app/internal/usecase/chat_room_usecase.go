package usecase

import (
	"app/internal/model"
	"app/internal/repository"
	"app/internal/validator"
)

type IChatRoomUsecase interface {
	Create(chatRoom model.ChatRoom) (model.ChatRoomResponse, error)
	Update(chatRoom model.ChatRoom) (model.ChatRoomResponse, error)
}

type chatRoomUsecase struct {
	crr repository.IChatRoomRepository
	crv validator.IChatRoomValidator
}

func NewChatRoomUsecase(crr repository.IChatRoomRepository, crv validator.IChatRoomValidator) IChatRoomUsecase {
	return &chatRoomUsecase{crr, crv}
}

func (cru *chatRoomUsecase) Create(chatRoom model.ChatRoom) (model.ChatRoomResponse, error) {
	if err := cru.crv.ChatRoomValidate(chatRoom); err != nil {
		return model.ChatRoomResponse{}, err
	}
	newChatRoom := model.ChatRoom{Title: chatRoom.Title, UserId: chatRoom.UserId}
	if err := cru.crr.CreateChatRoom(&newChatRoom); err != nil {
		return model.ChatRoomResponse{}, err
	}
	resChatRoom := model.ChatRoomResponse{
		ID:    newChatRoom.ID,
		Title: newChatRoom.Title,
	}
	return resChatRoom, nil
}

func (cru *chatRoomUsecase) Update(chatRoom model.ChatRoom) (model.ChatRoomResponse, error) {
	if err := cru.crv.ChatRoomValidate(chatRoom); err != nil {
		return model.ChatRoomResponse{}, err
	}
	updateChatRoom := model.ChatRoom{ID: chatRoom.ID, Title: chatRoom.Title, UserId: chatRoom.UserId}
	if err := cru.crr.UpdateChatRoom(&updateChatRoom); err != nil {
		return model.ChatRoomResponse{}, err
	}
	resChatRoom := model.ChatRoomResponse{
		ID:    updateChatRoom.ID,
		Title: updateChatRoom.Title,
	}
	return resChatRoom, nil
}

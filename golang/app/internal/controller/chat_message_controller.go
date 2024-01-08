package controller

import (
	"app/internal/model"
	"app/internal/usecase"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

type IChatMessageController interface {
	Create(c echo.Context) error
}

type chatMessageController struct {
	cmu usecase.IChatMessageUsecase
}

func NewChatMessageController(cmu usecase.IChatMessageUsecase) IChatMessageController {
	return &chatMessageController{cmu}
}

func (cmc *chatMessageController) Create(c echo.Context) error {
	chatMessage := model.ChatMessage{}
	if err := c.Bind(&chatMessage); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	val, err := strconv.ParseUint(c.Param("chat_room_id"), 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	chatMessage.ChatRoomId = uint(val)

	chatMessage.Answer, err = cmc.cmu.RequestChatGPTAnswer(chatMessage.ChatRoomId, chatMessage.Question)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	chatMessageRes, err := cmc.cmu.Create(chatMessage)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, chatMessageRes)
}

package controller

import (
	"app/internal/model"
	"app/internal/usecase"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

type IChatRoomController interface {
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type chatRoomController struct {
	cru usecase.IChatRoomUsecase
}

func NewChatRoomController(cru usecase.IChatRoomUsecase) IChatRoomController {
	return &chatRoomController{cru}
}

func (crc *chatRoomController) Create(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	chatRoom := model.ChatRoom{}
	if err := c.Bind(&chatRoom); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	chatRoom.UserId = uint(userId.(float64))
	chatRoomRes, err := crc.cru.Create(chatRoom)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, chatRoomRes)
}

func (crc *chatRoomController) Update(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	chatRoom := model.ChatRoom{}
	if err := c.Bind(&chatRoom); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	val, err := strconv.ParseUint(c.Param("chat_room_id"), 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	chatRoom.ID = uint(val)
	chatRoom.UserId = uint(userId.(float64))
	chatRoomRes, err := crc.cru.Create(chatRoom)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, chatRoomRes)
}

func (crc *chatRoomController) Delete(c echo.Context) error {
	return nil
}

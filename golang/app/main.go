package main

import (
	"app/internal/controller"
	"app/internal/db"
	"app/internal/repository"
	"app/internal/router"
	"app/internal/usecase"
	"app/internal/validator"
)

func main() {
	db := db.NewDB()
	userValidator := validator.NewUserValidator()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	userController := controller.NewUserController(userUsecase)

	chatRoomValidator := validator.NewChatRoomValidator()
	chatRoomRepository := repository.NewChatRoomRepository(db)
	chatRoomUsecase := usecase.NewChatRoomUsecase(chatRoomRepository, chatRoomValidator)
	chatRoomController := controller.NewChatRoomController(chatRoomUsecase)

	chatMessageValidator := validator.NewChatMessageValidator()
	chatMessageRepository := repository.NewChatMessageRepository(db)
	chatMessageUsecase := usecase.NewChatMessageUsecase(chatMessageRepository, chatMessageValidator)
	chatMessageController := controller.NewChatMessageController(chatMessageUsecase)

	e := router.NewRouter(userController, chatRoomController, chatMessageController)
	e.Logger.Fatal(e.Start(":8080"))
}

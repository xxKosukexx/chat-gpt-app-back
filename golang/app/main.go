package main

import (
	"app/internal/controller"
	"app/internal/db"
	"app/internal/pb"
	"app/internal/repository"
	"app/internal/router"
	"app/internal/usecase"
	"app/internal/validator"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

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

	go startGrpcServer()

	e.Logger.Fatal(e.Start(":8080"))
}

func startGrpcServer() {
	lis, err := net.Listen("tcp", ":50051") // gRPCポート
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &server{})
	fmt.Println("gRPC Server is running!")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (*server) GetUserEmails(ctx context.Context, in *pb.GetUserEmailsRequest) (*pb.GetUserEmailsResponse, error) {
	// TODO: DBからユーザーのメールアドレスを取得する
	fmt.Println("GetUserEmails is invoked!")
	return &pb.GetUserEmailsResponse{Emails: []string{"test@gmail.com", "test2@gmail.com"}}, nil
}

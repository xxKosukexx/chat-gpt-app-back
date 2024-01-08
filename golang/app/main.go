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
	userUsecase usecase.IUserUsecase
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

	server := &server{userUsecase: userUsecase}

	go startGrpcServer(server)

	e.Logger.Fatal(e.Start(":8080"))
}

func startGrpcServer(server *server) {
	lis, err := net.Listen("tcp", ":50051") // gRPCポート
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, server)
	fmt.Println("gRPC Server is running!")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) GetUserEmails(ctx context.Context, in *pb.GetUserEmailsRequest) (*pb.GetUserEmailsResponse, error) {
	fmt.Println("GetUserEmails is invoked!")
	emails, err := s.userUsecase.GetUserEmails()
	if err != nil {
		return nil, err
	}
	return &pb.GetUserEmailsResponse{Emails: emails}, nil
}

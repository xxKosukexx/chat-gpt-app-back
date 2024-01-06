package main

import (
	"context"
	"log"
	"time"

	"app/internal/pb" // プロトコルバッファパッケージへのパス
	"google.golang.org/grpc"
)

func main() {
	// gRPCサーバーへの接続を設定
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("サーバーへの接続に失敗: %v", err)
	}
	defer conn.Close()

	c := pb.NewUserServiceClient(conn)

	// タイムアウトの設定
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// サーバーにリクエストを送信
	r, err := c.GetUserEmails(ctx, &pb.GetUserEmailsRequest{})
	if err != nil {
		log.Fatalf("リクエストの実行に失敗: %v", err)
	}
	log.Printf("サーバーからのレスポンス: %s", r)
}

package main

import (
	"context"
	"log"
	"net"

	"github.com/kotokaze/ent-grpc-example/ent"
	"github.com/kotokaze/ent-grpc-example/ent/proto/entpb"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
)

func main() {
	// entクライアントの初期化
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	// マイグレーションツールの実行（テーブルの作成など）
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// 生成されたUserサービスの初期化
	svc := entpb.NewUserService(client)

	// 新しいgRPCサーバの作成 (複数のサービスを1つのサーバーに接続できます)
	server := grpc.NewServer()

	// Userサービスをサーバーに登録します
	entpb.RegisterUserServiceServer(server, svc)

	// トラフィックをリッスンするために5000番ポートを開きます
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed listening: %s", err)
	}

	// トラフィックを無期限にリッスン
	if err := server.Serve(lis); err != nil {
		log.Fatalf("server ended: %s", err)
	}
}

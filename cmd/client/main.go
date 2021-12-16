package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/kotokaze/ent-grpc-example/ent/proto/entpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// サーバーとのコネクションを開く
	conn, err := grpc.Dial(":5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed connecting to server: %s", err)
	}
	defer conn.Close()

	// コネクション上にUserサービスクライアントを作成
	client := entpb.NewUserServiceClient(conn)

	// サーバーにランダムなUserの作成を依頼する
	ctx := context.Background()
	user := randomUser()
	created, err := client.Create(ctx, &entpb.CreateUserRequest{
		User: user,
	})
	if err != nil {
		se, _ := status.FromError(err)
		log.Fatalf("failed creating user: status=%s message=%s", se.Code(), se.Message())
	}
	log.Printf("user created with id: %d", created.Id)

	// 別のRPC呼び出しで、先に作成したUserを取得
	get, err := client.Get(ctx, &entpb.GetUserRequest{
		Id: created.Id,
	})
	if err != nil {
		se, _ := status.FromError(err)
		log.Fatalf("failed retrieving user: status=%s message=%s", se.Code(), se.Message())
	}
	log.Printf("retrieved user with id=%d: %v", get.Id, get)
}

func randomUser() *entpb.User {
	return &entpb.User{
		Name:         fmt.Sprintf("user_%d", rand.Int()),
		EmailAddress: fmt.Sprintf("user_%d@example.com", rand.Int()),
		Alias:        wrapperspb.String("John Doe"),
	}
}

package main

import (
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/kotokaze/ent-grpc-example/ent/category"
	"github.com/kotokaze/ent-grpc-example/ent/enttest"
	"github.com/kotokaze/ent-grpc-example/ent/proto/entpb"
	"github.com/kotokaze/ent-grpc-example/ent/user"
)

func TestServiceWithEdges(t *testing.T) {
	// インメモリのsqliteインスタンスに接続されたentクライアントの初期化から始めます
	ctx := context.Background()
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	// 次に、Userサービスを初期化します。 ここでは、実際にポートを開いてgRPCサーバーを作成するのではなく
	// ライブラリのコードを直接呼び出していることに注目してください。
	svc := entpb.NewUserService(client)

	// 次に、entクライアントを使って直接Categoryを作成します。
	// Userとは無関係に初期化していることに注意してください。
	cat := client.Category.Create().SetName("cat_1").SaveX(ctx)

	// 次に、User サービスの `Create` メソッドを呼び出します。
	// IDのみが設定されたentpb.Categoryインスタンスのリストを渡していることに注意してください。
	create, err := svc.Create(ctx, &entpb.CreateUserRequest{
		User: &entpb.User{
			Name:         "user",
			EmailAddress: "user@service.code",
			Administered: []*entpb.Category{
				{Id: int32(cat.ID)},
			},
		},
	})
	if err != nil {
		t.Fatal("failed creating user using UserService", err)
	}

	// すべてが正しく動作したことを確認するために, カテゴリーテーブルをクエリします。
	// 作成したユーザーが管理するカテゴリーが1つだけあることを確認します。
	count, err := client.Category.
		Query().
		Where(
			category.HasAdminWith(
				user.ID(int(create.Id)),
			),
		).
		Count(ctx)
	if err != nil {
		t.Fatal("failed counting categories admin by created user", err)
	}
	if count != 1 {
		t.Fatal("expected exactly one group to managed by the created user")
	}
}

func TestGet(t *testing.T) {
	// インメモリのsqliteインスタンスに接続されたentクライアントの初期化から始めます
	ctx := context.Background()
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	// 次に、Userサービスを初期化します。 ここでは、実際にポートを開いてgRPCサーバーを作成するのではなく
	// ライブラリのコードを直接呼び出していることに注目してください。
	svc := entpb.NewUserService(client)

	// 次に、ユーザーとカテゴリーを作成し、そのユーザーをカテゴリーの管理者に設定します。
	user := client.User.Create().
		SetName("rotemtam").
		SetEmailAddress("r@entgo.io").
		SaveX(ctx)

	client.Category.Create().
		SetName("category").
		SetAdmin(user).
		SaveX(ctx)

	// 次に、エッジの情報なしでユーザーを取得します
	get, err := svc.Get(ctx, &entpb.GetUserRequest{
		Id: int32(user.ID),
	})
	if err != nil {
		t.Fatal("failed retrieving the created user", err)
	}
	if len(get.Administered) != 0 {
		t.Fatal("by default edge information is not supposed to be retrieved")
	}

	// 次に、エッジの情報*込み*でユーザーを取得します
	get, err = svc.Get(ctx, &entpb.GetUserRequest{
		Id:   int32(user.ID),
		View: entpb.GetUserRequest_WITH_EDGE_IDS,
	})
	if err != nil {
		t.Fatal("failed retrieving the created user", err)
	}
	if len(get.Administered) != 1 {
		t.Fatal("using WITH_EDGE_IDS edges should be returned")
	}
}

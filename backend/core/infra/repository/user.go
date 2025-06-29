package repository

import (
	"context"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/repository"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/db"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/dbmodels"
)

// user : pointerで保持
type user struct {
	dbClient *db.Client
}

// NewUser
// Userリポジトリのコンストラクタ
// 引数にはdb.Clientを受け取り、pointer型user構造体を返す
func NewUser(dbClient *db.Client) repository.User {
	return &user{dbClient}
}

// User作成のロジックの実装
// Createメソッドがポインタレシーバー（*user）で定義
// インターフェースrepository.Userも実装する必要がある
// ポインタレシーバーメソッドは、ポインタ型でないとインターフェースを満たせない
func (r *user) Create(ctx context.Context, user *model.User) (err error) {
	setter := &dbmodels.UserSetter{
		Ulid:      &user.Ulid,
		UID:       &user.UID,
		Nickname:  &user.Nickname,
		DeletedAt: &user.DeletedAt,
	}

	// Execによってbobがdb.Clientを呼び出し、データベースにユーザーを挿入する
	_, err = dbmodels.Users.Insert(setter).Exec(ctx, r.dbClient)
	if err != nil {
		return err
	}
	return nil
}

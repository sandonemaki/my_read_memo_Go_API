package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/oklog/ulid/v2"
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
	// ULIDが設定されていない場合は新しく生成
	if user.Ulid == "" {
		user.Ulid = ulid.Make().String()
	}

	setter := &dbmodels.UserSetter{
		Ulid:        &user.Ulid,
		UID:         &user.UID,
		DisplayName: &user.DisplayName,
		DeletedAt:   &user.DeletedAt,
	}

	// Execによってbobがdb.Clientを呼び出し、データベースにユーザーを挿入する
	_, err = dbmodels.Users.Insert(setter).Exec(ctx, r.dbClient)
	if err != nil {
		return err
	}
	return nil
}

func (r *user) Update(ctx context.Context, user *model.User) (ulid string, err error) {
	// 将来：複数条件が必要になった場合
	// modsをスライスに変更して条件を追加
	// var mods []bob.Mod[*dialect.UpdateQuery]
	// mods = append(mods, dbmodels.UpdateWhere.Users.Ulid.EQ(user.Ulid))
	// mods = append(mods, dbmodels.UpdateWhere.Users.DeletedAt.IsNull())
	setter := &dbmodels.UserSetter{
		DisplayName: &user.DisplayName,
	}

	mods := dbmodels.UpdateWhere.Users.Ulid.EQ(user.Ulid)
	_, err = dbmodels.Users.Update(setter.UpdateMod(), mods).Exec(ctx, r.dbClient)
	// _, err = dbmodels.Users.Update(setter.UpdateMod(), dbmodels.UpdateWhere.Users.Ulid.EQ(user.Ulid)).Exec(ctx, r.dbClient)
	if err != nil {
		return "", err
	}

	return user.Ulid, nil
}

// Deleteメソッドの実装
func (r *user) Delete(ctx context.Context, uid string) (err error) {
	mods := dbmodels.UpdateWhere.Users.UID.EQ(uid)
	now := time.Now()
	deletedTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	deletedAt := sql.Null[time.Time]{V: deletedTime, Valid: true}
	_, err = dbmodels.Users.Update(dbmodels.UserSetter{DeletedAt: &deletedAt}.UpdateMod(), mods).Exec(ctx, r.dbClient)
	if err != nil {
		return err
	}

	return nil
}

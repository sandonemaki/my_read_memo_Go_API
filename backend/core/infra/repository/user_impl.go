package repository

import (
	"context"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/repository"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/db"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/dbmodels"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/psql/dialect"
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

func (r *user) Get(ctx context.Context, query repository.UserGetQuery) (user *model.User, err error) {

	mods := []bob.Mod[*dialect.SelectQuery]{}

	if query.ULID.Valid {
		mods = append(mods, dbmodels.SelectWhere.Users.Ulid.EQ(query.ULID.String))
	}
	if query.UID.Valid {
		mods = append(mods, dbmodels.SelectWhere.Users.UID.EQ(query.UID.String))
	}

	dbUser, err := dbmodels.Users.Query(mods...).One(ctx, r.dbClient)
	if err != nil {
		return nil, err
	}

	return (*model.User)(dbUser), nil
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

func (r *user) Update(ctx context.Context, user *model.User) (ulid string, err error) {
	// 将来：複数条件が必要になった場合
	// modsをスライスに変更して条件を追加
	// var mods []bob.Mod[*dialect.UpdateQuery]
	// mods = append(mods, dbmodels.UpdateWhere.Users.Ulid.EQ(user.Ulid))
	// mods = append(mods, dbmodels.UpdateWhere.Users.DeletedAt.IsNull())
	setter := &dbmodels.UserSetter{
		Nickname: &user.Nickname,
	}

	mods := dbmodels.UpdateWhere.Users.Ulid.EQ(user.Ulid)
	_, err = dbmodels.Users.Update(setter.UpdateMod(), mods).Exec(ctx, r.dbClient)
	// _, err = dbmodels.Users.Update(setter.UpdateMod(), dbmodels.UpdateWhere.Users.Ulid.EQ(user.Ulid)).Exec(ctx, r.dbClient)
	if err != nil {
		return "", err
	}

	return user.Ulid, nil
}

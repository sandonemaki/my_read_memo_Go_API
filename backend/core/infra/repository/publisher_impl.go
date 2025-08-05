package repository

import (
	"context"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/repository"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/db"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/dbmodels"
)

type publisher struct {
	dbClient *db.Client
}

// NewPublisher creates a new publisher repository.
func NewPublisher(dbClient *db.Client) repository.Publisher {
	return &publisher{dbClient}
}

// Create creates a new publisher.
func (r *publisher) Create(ctx context.Context, publisher *model.Publisher) (publisherID int64, err error) {
	// Step 1: PublisherSetterを作成
	// Nameフィールドのみ設定（IDは自動生成、時刻はDB側で設定）
	setter := &dbmodels.PublisherSetter{
		Name: &publisher.Name,
	}

	// Step 2: Bob ORMでINSERT実行して作成されたレコードを取得
	// Insert()でINSERT文を構築、One()で実行して作成されたレコードを取得
	createdPublisher, err := dbmodels.Publishers.Insert(setter).One(ctx, r.dbClient.Get(ctx))
	if err != nil {
		return 0, err
	}

	// Step 3: 作成されたレコードのIDを返す
	return createdPublisher.ID, nil
}

// Update updates an existing publisher.
// TODO: OpenAPIにPUTエンドポイントを追加後に有効化
/*
func (r *publisher) Update(ctx context.Context, publisher *model.Publisher) (publisherID int64, err error) {
	// Step 1: PublisherSetterを作成
	setter := &dbmodels.PublisherSetter{
		Name: &publisher.Name,
	}

	// Step 2: WHERE句を作成（更新対象をIDで特定）
	mods := dbmodels.UpdateWhere.Publishers.ID.EQ(publisher.ID)

	// Step 3: UPDATE実行して更新されたレコードを取得
	// Update()でUPDATE文を構築、One()で実行して更新後のレコードを取得
	updatedPublisher, err := dbmodels.Publishers.Update(setter.UpdateMod(), mods).One(ctx, r.dbClient.Get(ctx))
	if err != nil {
		// sql.ErrNoRows の場合は更新対象が存在しない
		return 0, err
	}

	// Step 4: 更新されたレコードのIDを返す
	return updatedPublisher.ID, nil
}
*/
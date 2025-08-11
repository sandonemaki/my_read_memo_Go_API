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

func NewPublisher(dbClient *db.Client) repository.Publisher {
	return &publisher{dbClient}
}

func (r *publisher) Create(ctx context.Context, publisher *model.Publisher) (publisherID int64, err error) {
	setter := &dbmodels.PublisherSetter{
		Name: &publisher.Name,
	}

	// Bob ORMでINSERT実行して作成されたレコードを取得
	// Insert()でINSERT文を構築、One()で実行して作成されたレコードを取得
	createdPublisher, err := dbmodels.Publishers.Insert(setter).One(ctx, r.dbClient.Get(ctx))
	if err != nil {
		return 0, err
	}

	// 作成されたレコードのIDを返す
	return createdPublisher.ID, nil
}


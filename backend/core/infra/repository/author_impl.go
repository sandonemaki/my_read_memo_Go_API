package repository

import (
	"context"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/repository"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/db"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/dbmodels"
)

type author struct {
	dbClient *db.Client
}

// NewAuthor creates a new author repository.
func NewAuthor(dbClient *db.Client) repository.Author {
	return &author{dbClient}
}

// Create creates a new author.
func (r *author) Create(ctx context.Context, author *model.Author) (authorID int64, err error) {
	// Step 1: AuthorSetterを作成
	// Nameフィールドのみ設定（IDは自動生成、時刻はDB側で設定）
	setter := &dbmodels.AuthorSetter{
		Name: &author.Name,
	}
	
	// Step 2: Bob ORMでINSERT実行して作成されたレコードを取得
	// Insert()でINSERT文を構築、One()で実行して作成されたレコードを取得
	createdAuthor, err := dbmodels.Authors.Insert(setter).One(ctx, r.dbClient)
	if err != nil {
		return 0, err
	}
	
	// Step 3: 作成されたレコードのIDを返す
	return createdAuthor.ID, nil
}

// Update updates an existing author.
func (r *author) Update(ctx context.Context, author *model.Author) (authorID int64, err error) {
	// TODO: 一緒に実装しましょう
	return 0, nil
}

// Delete deletes an author by ID.
func (r *author) Delete(ctx context.Context, authorID int64) (err error) {
	// TODO: 一緒に実装しましょう
	return nil
}
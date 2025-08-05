package model

import (
	"database/sql"
	"time"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/dbmodels"
)

type MasterBook dbmodels.MasterBook

func NewMasterBook(
	isbn string,
	coverS3URL string,
	title string,
	authorID int64,
	publisherID int64,
	totalPage int32,
	createdAt time.Time,
	updatedAt time.Time,
	publishedAt sql.Null[time.Time],
) *MasterBook {
	return &MasterBook{
		Isbn:        isbn,
		CoverS3URL:  coverS3URL,
		Title:       title,
		AuthorID:    authorID,
		PublisherID: publisherID,
		TotalPage:   totalPage,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		PublishedAt: publishedAt,
	}
}

package input

import (
	"github.com/pkg/errors"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/errof"
)

func NewCreateAuthor(name string) CreateAuthor {
	return CreateAuthor{
		Name: name,
	}
}

type CreateAuthor struct {
	Name string `validate:"required,min=1,max=100"`
}

func (c *CreateAuthor) Validate() error {
	if err := validate.Struct(c); err != nil {
		return errors.Wrap(errof.ErrInvalidRequest, err.Error())
	}
	return nil
}

func NewGetAuthorByID(id int64) GetAuthorByID {
	return GetAuthorByID{
		ID: id,
	}
}

type GetAuthorByID struct {
	ID int64 `validate:"required,min=1"`
}

func (g *GetAuthorByID) Validate() error {
	if err := validate.Struct(g); err != nil {
		return errors.Wrap(errof.ErrInvalidRequest, err.Error())
	}
	return nil
}

func NewListAuthor() ListAuthor {
	return ListAuthor{}
}

type ListAuthor struct {
	// 現時点ではフィルター条件なし
	// 将来的にページングやソートが必要になったら追加
}

func NewSearchAuthor(name string) SearchAuthor {
	return SearchAuthor{
		Name: name,
	}
}

type SearchAuthor struct {
	Name string `validate:"required,min=1,max=100"`
}

func (s *SearchAuthor) Validate() error {
	if err := validate.Struct(s); err != nil {
		return errors.Wrap(errof.ErrInvalidRequest, err.Error())
	}
	return nil
}

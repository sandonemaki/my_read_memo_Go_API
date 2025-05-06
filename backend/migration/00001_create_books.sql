-- +goose Up
-- +goose StatementBegin
CREATE TABLE books (
  id BIGSERIAL PRIMARY KEY NOT NULL,
  title VARCHAR(60) NOT NULL DEFAULT '',
  author VARCHAR(60) NOT NULL DEFAULT '',
  publisher VARCHAR(60) DEFAULT '',
  total_page INTEGER NOT NULL DEFAULT 20,
  reading_state INTEGER NOT NULL DEFAULT 0,
  created_at timestamp DEFAULT statement_timestamp() NOT NULL,
  updated_at timestamp DEFAULT statement_timestamp() NOT NULL,
  seidoku_memo_key BOOLEAN DEFAULT TRUE,
  cover_path VARCHAR(255) DEFAULT '/illust/book_default_2d.png',
  user_id BIGINT NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE books;
-- +goose StatementEnd

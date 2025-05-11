-- +goose Up
-- +goose StatementBegin
CREATE TABLE seidoku_memos (
  id BIGSERIAL PRIMARY KEY NOT NULL,
  content TEXT DEFAULT '' NOT NULL,
  book_id BIGINT,
  content_state INTEGER DEFAULT 4 NOT NULL,
  created_at timestamp DEFAULT statement_timestamp() NOT NULL,
  updated_at timestamp DEFAULT statement_timestamp() NOT NULL,
  FOREIGN KEY (book_id) REFERENCES books(id)
);

CREATE TABLE seidoku_histories (
  id BIGSERIAL PRIMARY KEY NOT NULL,
  path VARCHAR(255) DEFAULT '' NOT NULL,
  book_id BIGINT DEFAULT 0,
  created_at timestamp DEFAULT statement_timestamp() NOT NULL,
  updated_at timestamp DEFAULT statement_timestamp() NOT NULL,
  FOREIGN KEY (book_id) REFERENCES books(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE seidoku_memos;
DROP TABLE seidoku_histories;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
CREATE TABLE randoku_histories (
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
DROP TABLE randoku_histories;
-- +goose StatementEnd

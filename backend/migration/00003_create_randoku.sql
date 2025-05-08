-- +goose Up
-- +goose StatementBegin
CREATE TABLE randoku_imgs (
  id BIGSERIAL PRIMARY KEY NOT NULL,
  first_post_flag INTEGER DEFAULT 0 NOT NULL,
  bookmark_flag INTEGER DEFAULT 0 NOT NULL,
  path VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  reading_state INTEGER DEFAULT 0 NOT NULL,
  book_id BIGINT,
  created_at timestamp DEFAULT statement_timestamp() NOT NULL,
  updated_at timestamp DEFAULT statement_timestamp() NOT NULL,
  thumbnail_path VARCHAR(255) NOT NULL,
  FOREIGN KEY (book_id) REFERENCES books(id)
);

CREATE TABLE randoku_memos (
  id BIGSERIAL PRIMARY KEY NOT NULL,
  content TEXT NOT NULL,
  content_state INTEGER DEFAULT 3 NOT NULL,
  book_id BIGINT,
  created_at timestamp DEFAULT statement_timestamp() NOT NULL,
  updated_at timestamp DEFAULT statement_timestamp() NOT NULL,
  FOREIGN KEY (book_id) REFERENCES books(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE randoku_imgs;
DROP TABLE randoku_memos;
-- +goose StatementEnd

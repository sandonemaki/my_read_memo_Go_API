-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  ulid TEXT PRIMARY KEY,
  nickname character varying(100) NOT NULL DEFAULT '',
  deleted_at timestamp DEFAULT NULL,
  created_at timestamp DEFAULT statement_timestamp() NOT NULL,
  updated_at timestamp DEFAULT statement_timestamp() NOT NULL
);

CREATE TABLE authors (
  id BIGSERIAL PRIMARY KEY NOT NULL,
  name VARCHAR(100) NOT NULL DEFAULT '',
  created_at timestamp DEFAULT statement_timestamp() NOT NULL,
  updated_at timestamp DEFAULT statement_timestamp() NOT NULL
);

CREATE TABLE publishers (
  id BIGSERIAL PRIMARY KEY NOT NULL,
  name VARCHAR(100) NOT NULL DEFAULT '',
  created_at timestamp DEFAULT statement_timestamp() NOT NULL,
  updated_at timestamp DEFAULT statement_timestamp() NOT NULL
);

CREATE TABLE master_books (
  id BIGSERIAL PRIMARY KEY NOT NULL,
  isbn VARCHAR(13) NOT NULL DEFAULT '',
  cover_s3_url VARCHAR(255) NOT NULL DEFAULT '',
  title VARCHAR(60) NOT NULL DEFAULT '',
  author_id BIGINT NOT NULL,
  publisher_id BIGINT NOT NULL,
  total_page INTEGER NOT NULL DEFAULT 20,
  created_at timestamp DEFAULT statement_timestamp() NOT NULL,
  updated_at timestamp DEFAULT statement_timestamp() NOT NULL,
  published_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
  FOREIGN KEY (author_id) REFERENCES authors(id),
  FOREIGN KEY (publisher_id) REFERENCES publishers(id)
);
CREATE TABLE user_book_logs (
  id BIGSERIAL PRIMARY KEY NOT NULL,
  user_ulid TEXT NOT NULL,
  master_book_id BIGINT NOT NULL,
  status INTEGER DEFAULT 0 NOT NULL,
  is_seidoku_key BOOLEAN DEFAULT FALSE NOT NULL,
  registered_at timestamp DEFAULT statement_timestamp() NOT NULL,
  created_at timestamp DEFAULT statement_timestamp() NOT NULL,
  UNIQUE (user_ulid, master_book_id),
  FOREIGN KEY (user_ulid) REFERENCES users(ulid),
  FOREIGN KEY (master_book_id) REFERENCES master_books(id)
);
CREATE TABLE randoku_images (
  ulid TEXT PRIMARY KEY,
  master_book_id BIGINT DEFAULT 0 NOT NULL,
  is_bookmark BOOLEAN DEFAULT FALSE NOT NULL,
  s3_url VARCHAR(255) NOT NULL DEFAULT '',
  thumbnail_s3_url VARCHAR(255) NOT NULL DEFAULT '',
  name VARCHAR(255) NOT NULL DEFAULT '',
  is_already_read BOOLEAN DEFAULT FALSE NOT NULL,
  created_at timestamp DEFAULT statement_timestamp() NOT NULL,
  updated_at timestamp DEFAULT statement_timestamp() NOT NULL,
  FOREIGN KEY (master_book_id) REFERENCES master_books(id)
);
CREATE TABLE randoku_memos (
  id BIGSERIAL PRIMARY KEY NOT NULL,
  master_book_id BIGINT NOT NULL,
  content TEXT DEFAULT '' NOT NULL,
  content_tag INTEGER DEFAULT 0 NOT NULL,
  created_at timestamp DEFAULT statement_timestamp() NOT NULL,
  updated_at timestamp DEFAULT statement_timestamp() NOT NULL,
  FOREIGN KEY (master_book_id) REFERENCES master_books(id)
);
CREATE TABLE seidoku_memos (
  id BIGSERIAL PRIMARY KEY NOT NULL,
  master_book_id BIGINT NOT NULL,
  content TEXT DEFAULT '' NOT NULL,
  content_tag INTEGER DEFAULT 0 NOT NULL,
  created_at timestamp DEFAULT statement_timestamp() NOT NULL,
  updated_at timestamp DEFAULT statement_timestamp() NOT NULL,
  FOREIGN KEY (master_book_id) REFERENCES master_books(id)
);
CREATE TABLE reading_history (
  id BIGSERIAL PRIMARY KEY NOT NULL,
  user_ulid TEXT NOT NULL,
  content_url VARCHAR(255) DEFAULT '' NOT NULL,
  recorded_at timestamp DEFAULT statement_timestamp() NOT NULL,
  created_at timestamp DEFAULT statement_timestamp() NOT NULL,
  FOREIGN KEY (user_ulid) REFERENCES users(ulid)
);
CREATE TABLE ocr_texts (
  id BIGSERIAL PRIMARY KEY NOT NULL,
  randoku_img_ulid TEXT NOT NULL,
  text VARCHAR(255) DEFAULT '' NOT NULL,
  created_at timestamp DEFAULT statement_timestamp() NOT NULL,
  updated_at timestamp DEFAULT statement_timestamp() NOT NULL,
  UNIQUE (randoku_img_ulid),
  FOREIGN KEY (randoku_img_ulid) REFERENCES randoku_images(ulid)
);
CREATE TABLE kindle_highlights (
  id BIGSERIAL PRIMARY KEY NOT NULL,
  master_book_id BIGINT NOT NULL,
  position INTEGER DEFAULT 0 NOT NULL,
  highlight TEXT DEFAULT '' NOT NULL,
  memo TEXT DEFAULT '' NOT NULL,
  last_synced_at timestamp DEFAULT statement_timestamp() NOT NULL,
  created_at timestamp DEFAULT statement_timestamp() NOT NULL,
  updated_at timestamp DEFAULT statement_timestamp() NOT NULL,
  UNIQUE (master_book_id, position),
  FOREIGN KEY (master_book_id) REFERENCES master_books(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
DROP TABLE master_books;
DROP TABLE authors;
DROP TABLE publishers;
DROP TABLE user_book_logs;
DROP TABLE randoku_images;
DROP TABLE randoku_memos;
DROP TABLE seidoku_memos;
DROP TABLE reading_history;
DROP TABLE ocr_texts;
DROP TABLE kindle_highlights;
-- +goose StatementEnd



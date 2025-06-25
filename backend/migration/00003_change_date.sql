-- +goose Up
-- +goose StatementBegin
ALTER TABLE master_books
ALTER COLUMN published_at DROP DEFAULT,
ALTER COLUMN published_at DROP NOT NULL,
ALTER COLUMN published_at TYPE date;

ALTER TABLE user_book_logs
ALTER COLUMN registered_at DROP DEFAULT,
ALTER COLUMN registered_at TYPE date;

ALTER TABLE kindle_highlights
ALTER COLUMN last_synced_at DROP DEFAULT,
ALTER COLUMN last_synced_at TYPE date;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE master_books
ALTER COLUMN published_at TYPE timestamp,
ALTER COLUMN published_at SET DEFAULT CURRENT_TIMESTAMP,
ALTER COLUMN published_at SET NOT NULL;

ALTER TABLE user_book_logs
ALTER COLUMN registered_at TYPE timestamp,
ALTER COLUMN registered_at SET DEFAULT statement_timestamp();

ALTER TABLE kindle_highlights
ALTER COLUMN last_synced_at TYPE timestamp;
ALTER COLUMN last_synced_at SET DEFAULT statement_timestamp();

-- +goose StatementEnd

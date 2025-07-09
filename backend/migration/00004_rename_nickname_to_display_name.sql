-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
  RENAME COLUMN nickname TO display_name;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
  RENAME COLUMN display_name TO nickname;
-- +goose StatementEnd

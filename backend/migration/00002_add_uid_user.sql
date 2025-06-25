-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
ADD COLUMN uid VARCHAR(255) NOT NULL DEFAULT '';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
DROP COLUMN uid;
-- +goose StatementEnd

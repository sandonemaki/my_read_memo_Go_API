-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  id BIGSERIAL PRIMARY KEY NOT NULL,
  uid character varying(100) NOT NULL,
  nickname character varying(100) NOT NULL,
  deactivated_at timestamp DEFAULT NULL,
  created_at timestamp DEFAULT statement_timestamp() NOT NULL,
  updated_at timestamp DEFAULT statement_timestamp() NOT NULL,
  CONSTRAINT users_uid_constraint UNIQUE (uid)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd

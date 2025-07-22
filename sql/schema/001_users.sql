-- +goose Up
-- use goose postgres "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable" up
-- and goose postgres "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable" down
-- to apply and revert this migration
CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE users;

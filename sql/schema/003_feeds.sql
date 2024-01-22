-- +goose Up
CREATE TABLE feeds (
    id UUID PRIMARY KEY,
    url TEXT UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'UTC'),
    updated_at TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'UTC')

);

-- +goose Down
DROP TABLE feeds;

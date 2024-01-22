-- name: CreateFeed :one
INSERT INTO feeds (id, url, name, user_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetUserFeeds :many
SELECT * FROM feeds WHERE user_id = $1;

-- name: GetAllFeeds :many
SELECT * FROM feeds;


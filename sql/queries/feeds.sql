-- name: CreateFeed :one
INSERT INTO feeds (id, url, name, user_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetUserFeeds :many
SELECT * FROM feeds WHERE user_id = $1;

-- name: GetAllFeeds :many
SELECT * FROM feeds;

-- name: GetNextFeedToFetch :many
SELECT * FROM feeds ORDER BY last_fetched_at ASC NULLS FIRST LIMIT $1;

-- name: UpdateLastFetch :one
UPDATE feeds SET last_fetched_at = NOW() WHERE id = $1 RETURNING *;


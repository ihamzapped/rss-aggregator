-- name: CreateFeedFollow :one
INSERT INTO feed_follows (id, user_id, feed_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetUserFollows :many
SELECT * FROM feed_follows WHERE user_id = $1;
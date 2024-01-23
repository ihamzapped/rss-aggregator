-- name: CreateFeedPost :exec
INSERT INTO feed_posts (id, title, description, url, feed_id, published_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
-- name: CreateFeedPost :exec
INSERT INTO feed_posts (id, title, description, url, feed_id, published_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetPostsForUser :many
SELECT feed_posts.* FROM feed_posts
JOIN feed_follows ON feed_posts.feed_id = feed_follows.feed_id
WHERE feed_follows.user_id = $1
ORDER BY feed_posts.published_at DESC
LIMIT $2;
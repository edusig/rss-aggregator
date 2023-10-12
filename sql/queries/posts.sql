-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
ON CONFLICT ON CONSTRAINT posts_url_key DO NOTHING
RETURNING *;

-- name: GetPostsByUser :many
SELECT posts.*
FROM posts
JOIN feeds ON feeds.id = posts.feed_id
JOIN users_feeds ON users_feeds.feed_id = feeds.id
WHERE users_feeds.user_id = $1
ORDER BY posts.published_at DESC, posts.updated_at DESC;
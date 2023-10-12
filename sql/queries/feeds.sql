-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetUserFeeds :many
SELECT * FROM feeds WHERE user_id=$1 LIMIT 100;

-- name: GetFeed :one
SELECT * FROM feeds WHERE id=$1;

-- name: GetAllFeeds :many
SELECT * FROM feeds LIMIT 100;

-- name: GetNextFeedsToFetch :many
SELECT url, id
FROM feeds
WHERE last_fetched_at IS NULL OR last_fetched_at < (NOW() - INTERVAL '2 hour')
ORDER BY last_fetched_at NULLS FIRST
LIMIT 10;

-- name: MarkFeedFetched :one
UPDATE feeds SET last_fetched_at=$1, updated_at=NOW() WHERE id=$2 RETURNING *;
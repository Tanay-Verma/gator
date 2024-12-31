-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
  uuid_generate_v4(),
  $1,
  $2,
  $3,
  $4,
  $5
)
RETURNING *;

-- name: GetFeeds :many
SELECT feeds.name AS feed_name, feeds.url, users.name AS user_name
FROM feeds
JOIN users ON feeds.user_id = users.id;

-- name: GetFeedByURL :one
SELECT *
FROM feeds
WHERE url = $1;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET updated_at = $2, last_fetched_at = $3
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT id, name, url
FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;

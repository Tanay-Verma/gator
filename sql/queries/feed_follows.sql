-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
  INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
  VALUES (
  uuid_generate_v4(),
  $1,
  $2,
  $3,
  $4
  )
  RETURNING *
)
SELECT
    inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
INNER JOIN feeds ON feeds.id = inserted_feed_follow.feed_id
INNER JOIN users ON users.id = inserted_feed_follow.user_id;

-- name: GetFeedFollowsForUser :many
SELECT
    feed_follows.id,
    feeds.name as feed_name,
    users.name as user_name
FROM feed_follows
JOIN users ON feed_follows.user_id = users.id
JOIN feeds ON feed_follows.feed_id = feeds.id
WHERE users.name = $1;


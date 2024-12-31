-- name: CreatePost :one
INSERT INTO posts (
  id,
  created_at,
  updated_at,
  title,
  url,
  description,
  published_at,
  feed_id
)
VALUES (uuid_generate_v4(), NOW(), NOW(), $1, $2, $3, $4, $5)
RETURNING *;


-- name: GetPostsForUser :many
SELECT posts.*, feeds.name AS feed_name
FROM posts
JOIN feed_follows ON feed_follows.feed_id = posts.feed_id
JOIN feeds ON feeds.id = posts.feed_id
WHERE feed_follows.user_id = $1
ORDER BY posts.published_at DESC
LIMIT $2;

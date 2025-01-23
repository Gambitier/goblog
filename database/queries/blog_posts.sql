-- name: CreateBlogPost :one
INSERT INTO blog_posts (title, description, body)
VALUES ($1, $2, $3)
RETURNING *;
-- name: GetBlogPost :one
SELECT *
FROM blog_posts
WHERE id = $1;
-- name: ListBlogPosts :many
SELECT *
FROM blog_posts
ORDER BY created_at DESC;
-- name: UpdateBlogPost :one
UPDATE blog_posts
SET title = $2,
    description = $3,
    body = $4
WHERE id = $1
RETURNING *;
-- name: DeleteBlogPost :exec
DELETE FROM blog_posts
WHERE id = $1;
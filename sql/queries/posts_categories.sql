
-- name: CreateCategory :one
INSERT INTO categories (name, slug) 
VALUES ($1, $2) 
RETURNING *;

-- name: ListCategories :many
SELECT * FROM categories ORDER BY name ASC;


-- name: CreatePost :one
INSERT INTO posts (title, slug, content, author_id, category_id, status)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetPostBySlug :one
SELECT * FROM posts 
WHERE slug = $1 AND status = 'published' LIMIT 1;

-- name: ListPostsByAuthor :many
SELECT * FROM posts 
WHERE author_id = $1 
ORDER BY created_at DESC;

-- name: UpdatePost :one
UPDATE posts 
SET title = $2, content = $3, status = $4, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts WHERE id = $1;

-- name: ListPublishedPosts :many
SELECT * FROM posts 
WHERE status = 'published' AND deleted_at IS NULL 
ORDER BY created_at DESC;

-- name: GetPostBySlugAdmin :one
SELECT * FROM posts 
WHERE slug = $1 AND status = 'published' AND deleted_at IS NULL LIMIT 1;

-- name: SoftDeletePost :exec
UPDATE posts 
SET deleted_at = NOW() 
WHERE id = $1;
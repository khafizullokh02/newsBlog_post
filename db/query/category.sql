-- name: CreateCategory :one
INSERT INTO category (
    title,
    post_count,
    article_count
)
VALUES (
    sqlc.arg(title),
    sqlc.arg(post_count),
    sqlc.arg(article_count)
)
RETURNING *;

-- name: GetCategory :one
SELECT *
FROM category
WHERE id = sqlc.arg(id)
LIMIT 1;

-- name: ListCategories :many
SELECT *
FROM category
ORDER BY id DESC
LIMIT sqlc.arg('limit')
OFFSET sqlc.arg('offset');

-- name: UpdateCategory :one
UPDATE category
SET
    title = COALESCE(sqlc.narg(title), title),
    post_count = COALESCE(sqlc.narg(post_count), post_count),
    article_count = COALESCE(sqlc.narg(article_count), article_count) 
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM category
WHERE id = sqlc.arg(id);
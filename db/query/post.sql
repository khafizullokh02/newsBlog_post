-- name: CreatePost :one
INSERT INTO post (
    user_id,
    title,
    content,
    category_id,
    post_type,
    like_count,
    comment_count,
    view_count,
    published_at,
    created_at,
    updated_at,
    deleted_at
)
VALUES (
    sqlc.arg(user_id),
    sqlc.arg(title),
    sqlc.arg(content),
    sqlc.arg(category_id),
    sqlc.arg(post_type),
    sqlc.arg(like_count),
    sqlc.arg(comment_count),
    sqlc.arg(view_count),
    sqlc.arg(published_at),
    sqlc.arg(created_at),
    sqlc.arg(updated_at),
    sqlc.arg(deleted_at)
)
RETURNING *;

-- name: GetPost :one
SELECT *
FROM post
WHERE id = sqlc.arg(id)
LIMIT 1;

-- name: ListPosts :many
SELECT *
FROM post
ORDER BY id DESC
LIMIT sqlc.arg('limit')
OFFSET sqlc.arg('offset');

-- name: UpdatePost :one
UPDATE post
SET
    user_id = COALESCE(sqlc.narg(user_id), user_id),
    title = COALESCE(sqlc.narg(title), title),
    content = COALESCE(sqlc.narg(content), content),
    category_id = COALESCE(sqlc.narg(category_id), category_id),
    post_type = COALESCE(sqlc.narg(post_type), post_type),
    like_count = COALESCE(sqlc.narg(like_count), like_count)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeletePost :exec
DELETE FROM post
WHERE id = sqlc.arg(id);
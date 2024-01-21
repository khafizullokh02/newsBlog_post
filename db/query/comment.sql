-- name: CreateComment :one
INSERT INTO comment (
    comment,
    parent_id,
    user_id,
    like_count,
    post_id
)
VALUES (
    sqlc.arg(comment),
    sqlc.arg(parent_id),
    sqlc.arg(user_id),
    sqlc.arg(like_count),
    sqlc.arg(post_id)
)
RETURNING *;

-- name: GetComment :one
SELECT *
FROM comment
WHERE id = sqlc.arg(id)
LIMIT 1;

-- name: ListComments :many
SELECT *
FROM comment
ORDER BY id DESC
LIMIT sqlc.arg('limit')
OFFSET sqlc.arg('offset');

-- name: UpdateComment :one
UPDATE comment
SET
    comment = COALESCE(sqlc.narg(comment), comment),
    parent_id = COALESCE(sqlc.arg(parent_id), parent_id),
    user_id = COALESCE(sqlc.narg(user_id), user_id),
    like_count = COALESCE(sqlc.narg(like_count), like_count),
    post_id = COALESCE(sqlc.narg(post_id), post_id)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteComment :exec
DELETE FROM comment
WHERE id = sqlc.arg(id);
-- name: CreateUser :one
INSERT INTO users (
    full_name,
    user_avatar,
    count_posts,
    email,
    pasword
)
VALUES (
    sqlc.arg(full_name),
    sqlc.arg(user_avatar),
    sqlc.arg(count_posts),
    sqlc.arg(email),
    sqlc.arg(pasword)
)
RETURNING *;

-- name: GetUser :one
SELECT *
FROM users
WHERE id = sqlc.arg(id)
LIMIT 1;

-- name: ListUsers :many
SELECT *
FROM users
WHERE full_name = sqlc.arg(full_name)
ORDER BY id DESC
LIMIT sqlc.arg('limit')
OFFSET sqlc.arg('offset');

-- name: UpdateUser :one
UPDATE users
SET
    full_name = COALESCE(sqlc.narg(full_name), full_name),
    user_avatar = COALESCE(sqlc.narg(user_avatar), user_avatar),
    count_posts = COALESCE(sqlc.narg(count_posts), count_posts),
    email = COALESCE(sqlc.narg(email), email),
    pasword = COALESCE(sqlc.narg(pasword), pasword)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = sqlc.arg(id);
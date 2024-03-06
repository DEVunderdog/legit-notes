-- name: CreateUser :one
INSERT INTO users (
    username,
    email,
    password
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM users
WHERE id = $1;
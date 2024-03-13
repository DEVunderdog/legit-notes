-- name: CreateUser :one
INSERT INTO users (
    username,
    email,
    hashed_password
) VALUES (
    $1, $2, $3
) RETURNING *;


-- name: DeleteAccount :exec
DELETE FROM users
WHERE id = $1;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;
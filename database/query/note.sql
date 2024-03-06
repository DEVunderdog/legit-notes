-- name: CreateNote :one
INSERT INTO notes (
    user_id,
    title,
    description
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetNote :one
SELECT * FROM notes
WHERE id = $1 LIMIT 1;

-- name: ListNotes :many
SELECT * FROM notes
WHERE user_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: DeleteNote :exec
DELETE FROM notes
WHERE id = $1;

-- name: DeleteAllNotes :exec
DELETE FROM "notes";
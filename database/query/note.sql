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

-- name: UpdateNoteTitle :one
UPDATE notes SET title = COALESCE($2, title),
updated_at = now()
WHERE id = $1 RETURNING *;

-- name: UpdateNoteDescription :one
UPDATE notes SET description = COALESCE($2, description),
updated_at = now()
WHERE id = $1 RETURNING *;

-- name: UpdateNote :one
UPDATE notes SET title = COALESCE($2, title),
description = COALESCE($3, description),
updated_at = now()
WHERE id = $1 RETURNING *;

-- name: DeleteNote :exec
DELETE FROM notes
WHERE id = $1;

-- name: DeleteAllNotes :exec
DELETE FROM "notes";
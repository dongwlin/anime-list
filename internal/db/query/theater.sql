-- name: CreateTheater :one
INSERT INTO theaters (
    anime_id,
    name,
    cover,
    released_at,
    description,
    status
) VALUES (
    ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: GetTheater :one
SELECT * FROM theaters
WHERE id = ?
LIMIT 1;

-- name: ListTheater :many
SELECT * FROM theaters
WHERE anime_id = ?
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: CountTheater :one
SELECT COUNT(*) as total
FROM theaters
WHERE anime_id = ?;

-- name: UpdateTheater :one
UPDATE theaters
SET name = ?, cover = ?, released_at = ?, description = ?, status = ?
WHERE id = ?
RETURNING *;

-- name: DeleteTheater :exec
DELETE FROM theaters
WHERE id = ?;
-- name: CreateTheater :one
INSERT INTO theaters (
    anime_id,
    name,
    cover,
    released_at,
    desc,
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
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: UpdateTheater :one
UPDATE theaters
SET name = ?, cover = ?, released_at = ?, desc = ?, status = ?, updated_at = strftime('%Y-%m-%d %H:%M:%f', 'now', 'localtime')
WHERE id = ?
RETURNING *;

-- name: DeleteTheater :exec
DELETE FROM theaters
WHERE id = ?;
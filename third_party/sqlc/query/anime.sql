-- name: CreateAnime :one
INSERT INTO animes (
    name,
    desc,
    status
) VALUES (
    ?, ?, ?
)
RETURNING *;

-- name: GetAnime :one
SELECT * FROM animes
WHERE id = ?
LIMIT 1;

-- name: ListAnime :many
SELECT * FROM animes
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: UpdateAnime :one
UPDATE animes
SET name = ?, desc = ?, status = ?, updated_at = strftime('%Y-%m-%d %H:%M:%f', 'now', 'localtime')
WHERE id = ?
RETURNING *;

-- name: DeleteAnime :exec
DELETE FROM animes
WHERE id = ?;
-- name: CreateAnime :one
INSERT INTO animes (
    name,
    description,
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

-- name: CountAnime :one
SELECT COUNT(*) AS total
FROM animes;

-- name: UpdateAnime :one
UPDATE animes
SET name = ?, description = ?, status = ?
WHERE id = ?
RETURNING *;

-- name: DeleteAnime :exec
DELETE FROM animes
WHERE id = ?;
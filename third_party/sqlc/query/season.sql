-- name: CreateSeason :one
INSERT INTO seasons (
    anime_id,
    name,
    value,
    cover,
    released_at,
    desc,
    status
) VALUES (
    ?, ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: GetSeason :one
SELECT * FROM seasons
WHERE id = ?
LIMIT 1;

-- name: ListSeason :many
SELECT * FROM seasons
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: UpdateSeason :one
UPDATE seasons
SET name = ?, value = ?, cover = ?, released_at = ?, desc = ?, status = ?, updated_at = strftime('%Y-%m-%d %H:%M:%f', 'now', 'localtime')
WHERE id = ?
RETURNING *;

-- name: DeleteSeason :exec
DELETE FROM seasons
WHERE id = ?;
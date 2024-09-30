-- name: CreateSeason :one
INSERT INTO seasons (
    anime_id,
    name,
    value,
    cover,
    released_at,
    description,
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
WHERE anime_id = ?
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: CountSeason :one
SELECT COUNT(*) as total
FROM seasons
WHERE anime_id = ?;

-- name: UpdateSeason :one
UPDATE seasons
SET name = ?, value = ?, cover = ?, released_at = ?, description = ?, status = ?, updated_at = strftime('%Y-%m-%d %H:%M:%f', 'now', 'localtime')
WHERE id = ?
RETURNING *;

-- name: DeleteSeason :exec
DELETE FROM seasons
WHERE id = ?;
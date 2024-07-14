-- name: CreateEpisode :one
INSERT INTO episodes (
    season_id,
    name,
    value,
    desc,
    status
) VALUES (
    ?, ?, ?, ?, ?
)
RETURNING *;

-- name: GetEpisode :one
SELECT * FROM episodes
WHERE id = ?
LIMIT 1;

-- name: ListEpisode :many
SELECT * FROM episodes
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: UpdateEpisode :one
UPDATE episodes
SET name = ?, value = ?, desc = ?, status = ?, updated_at = strftime('%Y-%m-%d %H:%M:%f', 'now', 'localtime')
WHERE id = ?
RETURNING *;

-- name: DeleteEpisode :exec
DELETE FROM episodes
WHERE id = ?;
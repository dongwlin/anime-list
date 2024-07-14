-- name: CreateUser :one
INSERT INTO users (
    username,
    hashed_password,
    is_admin,
    desc,
    status
) VALUES (
    ?, ?, ?, ?, ?
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = ?
LIMIT 1;

-- name: ListUser :many
SELECT * FROM users
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: UpdateUser :one
UPDATE users
SET hashed_password = ?, is_admin = ?, desc = ?, status = ?, updated_at = strftime('%Y-%m-%d %H:%M:%f', 'now', 'localtime')
WHERE id = ?
RETURNING *;

-- name: ChangePassword :one
UPDATE users
SET hashed_password = ?, updated_at = strftime('%Y-%m-%d %H:%M:%f', 'now')
WHERE id = ?
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;
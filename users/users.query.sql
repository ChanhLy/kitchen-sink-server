-- name: GetUserById :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users;


-- name: CreateUser :one
INSERT INTO users (name, email) VALUES (?, ?) RETURNING *;

-- name: UpdateUser :exec
UPDATE users SET name = ? where id = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;

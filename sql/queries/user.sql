-- name: CreateUser :one
INSERT INTO users (email, password, name) VALUES ($1, $2, $3) RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: UpdateUser :one
UPDATE users SET email = $2, password = $3, name = $4 WHERE id = $1 RETURNING *;
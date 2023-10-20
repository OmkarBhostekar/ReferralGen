-- name: CreateUser :one
INSERT INTO users(
    username,
    hashed_password,
    email
) VALUES (
    $1,
    $2,
    $3
) RETURNING *;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;
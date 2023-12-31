// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: users.sql

package referralgen

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users(
    fullname,
    email,
    hashed_password
) VALUES (
    $1,
    $2,
    $3
) RETURNING id, fullname, hashed_password, email, created_at, password_changed_at
`

type CreateUserParams struct {
	Fullname       string `json:"fullname"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Fullname, arg.Email, arg.HashedPassword)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Fullname,
		&i.HashedPassword,
		&i.Email,
		&i.CreatedAt,
		&i.PasswordChangedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, fullname, hashed_password, email, created_at, password_changed_at FROM users WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Fullname,
		&i.HashedPassword,
		&i.Email,
		&i.CreatedAt,
		&i.PasswordChangedAt,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, fullname, hashed_password, email, created_at, password_changed_at FROM users WHERE id = $1
`

func (q *Queries) GetUserById(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Fullname,
		&i.HashedPassword,
		&i.Email,
		&i.CreatedAt,
		&i.PasswordChangedAt,
	)
	return i, err
}

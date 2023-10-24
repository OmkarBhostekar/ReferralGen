// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: generations.sql

package referralgen

import (
	"context"
)

const createGeneration = `-- name: CreateGeneration :one
INSERT INTO generations (user_id,created_date) VALUES ($1, $2) RETURNING id, user_id, created_date, count
`

type CreateGenerationParams struct {
	UserID      int64  `json:"user_id"`
	CreatedDate string `json:"created_date"`
}

func (q *Queries) CreateGeneration(ctx context.Context, arg CreateGenerationParams) (Generation, error) {
	row := q.db.QueryRowContext(ctx, createGeneration, arg.UserID, arg.CreatedDate)
	var i Generation
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CreatedDate,
		&i.Count,
	)
	return i, err
}

const getTodayUserGenerationCount = `-- name: GetTodayUserGenerationCount :one
select id, user_id, created_date, count from generations where user_id = $1 and created_date = $2
`

type GetTodayUserGenerationCountParams struct {
	UserID      int64  `json:"user_id"`
	CreatedDate string `json:"created_date"`
}

func (q *Queries) GetTodayUserGenerationCount(ctx context.Context, arg GetTodayUserGenerationCountParams) (Generation, error) {
	row := q.db.QueryRowContext(ctx, getTodayUserGenerationCount, arg.UserID, arg.CreatedDate)
	var i Generation
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CreatedDate,
		&i.Count,
	)
	return i, err
}

const increaseGenerationCount = `-- name: IncreaseGenerationCount :one
UPDATE generations SET count = count + 1 WHERE id = $1 RETURNING id, user_id, created_date, count
`

func (q *Queries) IncreaseGenerationCount(ctx context.Context, id int64) (Generation, error) {
	row := q.db.QueryRowContext(ctx, increaseGenerationCount, id)
	var i Generation
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CreatedDate,
		&i.Count,
	)
	return i, err
}

-- name: GetTodayUserGenerationCount :one
select * from generations where user_id = $1 and created_date = $2;

-- name: CreateGeneration :one
INSERT INTO generations (user_id,created_date) VALUES ($1, $2) RETURNING *;

-- name: IncreaseGenerationCount :one
UPDATE generations SET count = count + 1 WHERE id = $1 RETURNING *;
-- name: CreateTemplate :one
INSERT INTO templates(
    user_id,
    name,
    template,
    params
) VALUES (
    $1,
    $2,
    $3,
    $4
) RETURNING *;

-- name: GetTemplateById :one
SELECT * FROM templates WHERE id = $1;

-- name: GetTemplatesByUser :many
SELECT * FROM templates WHERE user_id = $1;

-- name: GetTemplateByName :one
SELECT * FROM templates WHERE user_id = $1 AND name = $2;

-- name: DeleteTemplateById :one
DELETE FROM templates WHERE id = $1 RETURNING *;

-- name: UpdateTemplateById :one
UPDATE templates SET
    name = $2,
    template = $3,
    params = $4
WHERE id = $1 RETURNING *;


-- name: CreateLanguage :one
INSERT INTO languages (code, name)
VALUES (
    $1,
    $2
)
RETURNING *;
-- name: CreateWord :one
INSERT INTO words (token_name, last_seen_at, known, frequency)
VALUES (
    $1,
    now(),
    $2,
    1
)
RETURNING *;
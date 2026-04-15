-- name: CreateWord :one
INSERT INTO words (token_name, last_seen_at, known, frequency, language_id)
VALUES (
    $1,
    now(),
    $2,
    1,
    $3
)
RETURNING *;
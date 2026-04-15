-- name: GetWord :one
SELECT * FROM words WHERE token_name = $1 AND language_id = $2;
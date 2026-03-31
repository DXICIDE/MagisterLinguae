-- name: GetWord :one
SELECT * FROM words WHERE token_name = $1;
-- name: GetWords :many
SELECT * FROM words
WHERE language_id = $1;
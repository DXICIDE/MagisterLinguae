-- name: ResetLangWords :exec
DELETE FROM words
WHERE language_id = $1;

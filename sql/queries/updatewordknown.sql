-- name: UpdateWordKnown :exec
UPDATE words
SET known = true
WHERE token_name = $1 AND language_id = $2;

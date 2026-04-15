-- name: UpdateWordUnKnown :exec
UPDATE words
SET known = false
WHERE token_name = $1 AND language_id = $2;

-- name: UpdateWordKnown :exec
UPDATE words
SET known = true
WHERE token_name = $1;

-- name: UpdateWordPrompted :exec
UPDATE words
SET promted_user_to_mark = true
WHERE token_name = $1 AND language_id = $2;

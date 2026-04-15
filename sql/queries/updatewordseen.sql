-- name: UpdateWordSeen :exec
UPDATE words
SET last_seen_at = now(), frequency = frequency + 1
WHERE token_name = $1 AND language_id = $2;

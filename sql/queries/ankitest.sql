-- name: GetAnki :many
SELECT * FROM words
WHERE frequency > 3 AND language_id = $1 AND known = true
ORDER BY last_seen_at ASC 
LIMIT 30;
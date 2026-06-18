-- name: GetAnki :many
SELECT * FROM words
WHERE frequency > 1 AND language_id = $1 AND known = true AND last_seen_at < NOW() - INTERVAL '7 days'
ORDER BY last_seen_at ASC 
LIMIT $2;
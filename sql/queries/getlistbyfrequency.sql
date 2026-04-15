-- name: GetListByFrequency :many
SELECT * FROM words WHERE frequency > 1 AND language_id = $1 ORDER BY known ASC, frequency DESC;
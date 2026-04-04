-- name: MarkWordsByFrequency :many
SELECT * FROM words WHERE frequency = 7 AND known = false AND promted_user_to_mark = false;
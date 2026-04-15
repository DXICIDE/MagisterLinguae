-- name: GetLanguageByCode :one
SELECT * FROM languages WHERE code = $1;
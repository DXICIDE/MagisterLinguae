-- name: GetLanguageById :one
SELECT * FROM languages WHERE id = $1;
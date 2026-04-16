-- name: DeleteLanguage :exec
DELETE FROM languages WHERE code = $1;
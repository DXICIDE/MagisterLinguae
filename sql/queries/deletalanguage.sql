-- name: DeleteLanguage :exec
DELETE FROM languages WHERE id = $1;
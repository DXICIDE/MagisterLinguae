-- name: ResetWords :exec
TRUNCATE words RESTART IDENTITY CASCADE;

-- name: ResetLanguages :exec
TRUNCATE languages RESTART IDENTITY CASCADE;
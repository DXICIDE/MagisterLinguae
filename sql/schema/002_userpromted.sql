-- +goose Up
ALTER TABLE words
ADD promted_user_to_mark BOOLEAN NOT NULL DEFAULT 'false';

-- +goose Down
ALTER TABLE words
DROP COLUMN promted_user_to_mark;
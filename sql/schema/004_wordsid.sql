-- +goose Up
ALTER TABLE words
ADD language_id INT NOT NULL REFERENCES languages(id);

-- +goose Down
ALTER TABLE words
DROP COLUMN language_id;
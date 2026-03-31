-- +goose Up
CREATE TABLE words (
token_name TEXT PRIMARY KEY,
last_seen_at TIMESTAMP NOT NULL,
known BOOLEAN NOT NULL,
frequency INT NOT NULL);

-- +goose Down
DROP TABLE words;
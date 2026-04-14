-- +goose Up
CREATE TABLE languages (
id SERIAL PRIMARY KEY,
code VARCHAR(10) UNIQUE,
name VARCHAR(50));

-- +goose Down
DROP TABLE languages;
-- +goose Up
CREATE TABLE languages (
id SERIAL PRIMARY KEY,
code VARCHAR(10) NOT NULL,
name VARCHAR(50) NOT NULL);

-- +goose Down
DROP TABLE languages;
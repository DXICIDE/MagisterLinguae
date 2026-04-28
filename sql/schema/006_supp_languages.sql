-- +goose Up
INSERT INTO languages (code, name) VALUES ('it', 'Italian');
INSERT INTO languages (code, name) VALUES ('cz', 'Czech');

-- +goose Down
DELETE FROM languages WHERE id = 1;
DELETE FROM languages WHERE id = 2;
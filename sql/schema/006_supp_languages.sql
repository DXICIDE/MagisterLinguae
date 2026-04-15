-- +goose Up
INSERT INTO languages (code, name) VALUES ('it', 'Italian');
INSERT INTO languages (code, name) VALUES ('cz', 'Czech');

-- +goose Down
DELETE FROM languages WHERE code = 'it';
DELETE FROM languages WHERE code = 'cz';
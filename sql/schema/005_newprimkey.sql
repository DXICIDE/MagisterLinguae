-- +goose Up
ALTER TABLE words DROP CONSTRAINT words_pkey;
ALTER TABLE words ADD PRIMARY KEY (token_name, language_id);

-- +goose Down
ALTER TABLE words DROP CONSTRAINT words_pkey;
ALTER TABLE words ADD PRIMARY KEY (token_name);
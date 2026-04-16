-- name: SeedLanguages :exec
INSERT INTO languages (code, name) VALUES 
    ('it', 'Italian'),
    ('cz', 'Czech')
ON CONFLICT (code) DO NOTHING;
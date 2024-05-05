-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE IF NOT EXISTS users(
    id bigserial PRIMARY KEY,
    created_at timestamp(0) NOT NULL DEFAULT NOW(),
    github_id bigint,
    name text,
    email text UNIQUE,
    image text, 
    password text, 
    verified boolean NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS tokens (
    hash bytea PRIMARY KEY,
    user_id bigint NOT NULL REFERENCES users ON DELETE CASCADE,
    expiry timestamp NOT NULL DEFAULT NOW(),
    scope text NOT NULL
);

CREATE INDEX IF NOT EXISTS users_email_idx ON users USING GIN (to_tsvector('simple', email));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS tokens;
DROP INDEX IF EXISTS users_email_idx;

-- +goose StatementEnd

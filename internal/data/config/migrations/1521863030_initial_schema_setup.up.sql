BEGIN;

CREATE TABLE IF NOT EXISTS sample (
    id serial PRIMARY KEY,
    text text NOT NULL
);

COMMIT;
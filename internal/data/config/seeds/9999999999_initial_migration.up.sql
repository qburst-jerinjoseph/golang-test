DROP TABLE IF EXISTS schema_migrations;
CREATE TABLE IF NOT EXISTS schema_migrations (
    version bigint PRIMARY KEY,
    dirty boolean NOT NULL
);
DELETE FROM schema_migrations;
INSERT INTO schema_migrations (version, dirty) VALUES(__VER__, false);
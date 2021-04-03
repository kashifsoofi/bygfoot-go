CREATE TABLE IF NOT EXISTS association (
    id integer NOT NULL PRIMARY KEY AUTOINCREMENT,
    sid TEXT NOT NULL,
    name TEXT NOT NULL,
    symbol TEXT NULL,
    parent_id TEXT NULL
);
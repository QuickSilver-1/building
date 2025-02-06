CREATE TABLE building (
    id      SERIAL PRIMARY KEY,
    name    VARCHAR(256) NOT NULL,
    city    VARCHAR(64) NOT NULL,
    year    INTEGER,
    floor   INTEGER
);
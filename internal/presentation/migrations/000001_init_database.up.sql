CREATE TABLE building (
    id      SERIAL PRIMARY KEY,
    name    VARCHAR(256) NOT NULL,
    city    VARCHAR(64) NOT NULL,
    year    INTEGER,
    floor   INTEGER

    CREATE INDEX idx_city ON building(city);
    CREATE INDEX idx_year ON building(year);
    CREATE INDEX idx_floor ON building(floor);
);
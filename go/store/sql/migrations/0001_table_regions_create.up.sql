CREATE TABLE IF NOT EXISTS Regions (
    Id          INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    Sid         TEXT NOT NULL,
    Name        TEXT NOT NULL,
    Symbol      TEXT NULL,
    Rating      INTEGER NOT NULL,
    RegionType  INTEGER NOT NULL,
    ParentId   INTEGER NULL
);

INSERT INTO Regions
    (Sid, Name, Symbol, Rating, RegionType, ParentId)
VALUES
    ('world', 'World', '', 10, 0, NULL),
    ('africa', 'Africa', '', 9, 1, 1),
    ('asia', 'Asia', '', 9, 1, 1),
    ('australia', 'Australia', '', 9, 1, 1),
    ('europe', 'Europe', '', 10, 1, 1),
    ('north_america', 'North America', '', 9, 1, 1),
    ('south_america', 'South America', '', 10, 1, 1),
    ('algeria', 'Algeria', 'flag_algeria.png', 9, 2, 2),
    ('england', 'England', 'flag_england.png', 10, 2, 5),
    ('france', 'France', 'flag_france.png', 9, 2, 5)

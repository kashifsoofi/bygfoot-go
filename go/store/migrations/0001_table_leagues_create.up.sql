CREATE TABLE IF NOT EXISTS Leagues (
    Id              INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    Sid             TEXT NOT NULL,
    Name            TEXT NOT NULL,
    ShortName       TEXT NOT NULL,
    Symbol          TEXT NULL,
    PromRel         INTEGER NOT NULL,
    Layer           INTEGER NOT NULL,
    FirstWeek       INTEGER NOT NULL,
    WeekGap         INTEGER NOT NULL,
    RoundRobins     INTEGER NOT NULL,
    RrBreak         INTEGER NOT NULL,
    YellowRed       INTEGER NOT NULL,
    AverageTalent   FLOAT NOT NULL,
    Active          BIT NOT NULL,
    RegionId        INTEGER NOT NULL,
    FOREIGN KEY(RegionId) REFERENCES Regions(Id)
);

INSERT INTO Leagues
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

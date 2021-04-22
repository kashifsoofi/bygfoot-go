CREATE TABLE IF NOT EXISTS Leagues (
    Id              INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    Sid             TEXT NOT NULL,
    Name            TEXT NOT NULL,
    ShortName       TEXT NOT NULL,
    Symbol          TEXT NULL,
    FirstWeek       INTEGER NOT NULL,
    WeekGap         INTEGER NOT NULL,
    YellowRed       INTEGER NOT NULL,
    AverageTalent   FLOAT NOT NULL,
    Active          BIT NOT NULL,
    RegionId        INTEGER NOT NULL,
    FOREIGN KEY(RegionId) REFERENCES Regions(Id)
);

INSERT INTO Leagues
    (Sid, Name, ShortName, Symbol, FirstWeek, WeekGap, YellowRed, AverageTalent, Active, RegionId)
VALUES
    ('england1', 'Premier League', 'EPL', 'premier.png', 7, 1, 5, 8100, 1, 9),
    ('england2', 'Football League Championship', 'Championship', 'championship.png', 6, 1, 5, 6600, 1, 9),
    ('england3', 'Football League One', 'League One', 'championship.png', 6, 1, 5, 5800, 1, 9),
    ('england4', 'Football League Two', 'League Two', 'championship.png', 6, 1, 5, 5200, 1, 9),
    ('england5', 'National Conference', 'Conference', 'flag_england.png', 6, 1, 5, 4800, 1, 9),
    ('france1', 'Ligue 1', 'L1', 'ligue1.png', 6, 1, 5, 8100, 1, 10),
    ('france2', 'Ligue 2', 'L2', 'flag_france.png', 6, 1, 5, 6300, 1, 10),
    ('france3', 'Championnat de France National', 'Ntl', 'flag_france.png', 6, 1, 5, 5700, 1, 10)

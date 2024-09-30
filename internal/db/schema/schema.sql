-- table
CREATE TABLE IF NOT EXISTS animes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    status INTEGER NOT NULL,
    created_at DATETIME NOT NULL DEFAULT (strftime('%Y-%m-%d %H:%M:%f', 'now', 'localtime')),
    updated_at DATETIME NOT NULL DEFAULT (strftime('%Y-%m-%d %H:%M:%f', 'now', 'localtime'))
);

CREATE TABLE IF NOT EXISTS seasons (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    anime_id INTEGER NOT NULL,
    name VARCHAR(255) NOT NULL,
    value INTEGER NOT NULL,
    cover TEXT NOT NULL,
    released_at DATETIME NOT NULL,
    description TEXT NOT NULL,
    status INTEGER NOT NULL,
    created_at DATETIME NOT NULL DEFAULT (strftime('%Y-%m-%d %H:%M:%f', 'now', 'localtime')),
    updated_at DATETIME NOT NULL DEFAULT (strftime('%Y-%m-%d %H:%M:%f', 'now', 'localtime')),
    FOREIGN KEY (anime_id) REFERENCES animes(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS episodes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    season_id INTEGER NOT NULL,
    name VARCHAR(255) NOT NULL,
    value INTEGER NOT NULL,
    description TEXT NOT NULL,
    status INTEGER NOT NULL,
    created_at DATETIME NOT NULL DEFAULT (strftime('%Y-%m-%d %H:%M:%f', 'now', 'localtime')),
    updated_at DATETIME NOT NULL DEFAULT (strftime('%Y-%m-%d %H:%M:%f', 'now', 'localtime')),
    FOREIGN KEY (season_id) REFERENCES seasons(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS theaters (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    anime_id INTEGER NOT NULL,
    name VARCHAR(255) NOT NULL,
    cover TEXT NOT NULL,
    released_at DATETIME NOT NULL,
    description TEXT NOT NULL,
    status INTEGER NOT NULL,
    created_at DATETIME NOT NULL DEFAULT (strftime('%Y-%m-%d %H:%M:%f', 'now', 'localtime')),
    updated_at DATETIME NOT NULL DEFAULT (strftime('%Y-%m-%d %H:%M:%f', 'now', 'localtime')),
    FOREIGN KEY (anime_id) REFERENCES animes(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- index
CREATE INDEX idx_anime_id ON seasons(anime_id);
CREATE INDEX idx_season_id ON episodes(season_id);
CREATE INDEX idx_anime_id_theaters ON theaters(anime_id);

-- trigger
CREATE TRIGGER update_anime_timestamp
    AFTER UPDATE ON animes
    FOR EACH ROW
BEGIN
    UPDATE animes SET updated_at = strftime('%Y-%m-%d %H:%M:%f', 'now', 'localtime') WHERE id = OLD.id;
END;

CREATE TRIGGER update_season_timestamp
    AFTER UPDATE ON seasons
    FOR EACH ROW
BEGIN
    UPDATE seasons SET updated_at = strftime('%Y-%m-%d %H:%M:%f', 'now', 'localtime') WHERE id = OLD.id;
END;

CREATE TRIGGER update_episode_timestamp
    AFTER UPDATE ON episodes
    FOR EACH ROW
BEGIN
    UPDATE episodes SET updated_at = strftime('%Y-%m-%d %H:%M:%f', 'now', 'localtime') WHERE id = OLD.id;
END;

CREATE TRIGGER update_theater_timestamp
    AFTER UPDATE ON theaters
    FOR EACH ROW
BEGIN
    UPDATE theaters SET updated_at = strftime('%Y-%m-%d %H:%M:%f', 'now', 'localtime') WHERE id = OLD.id;
END;

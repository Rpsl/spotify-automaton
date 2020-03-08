-- name: drop-table-tracks
DROP TABLE IF EXISTS tracks;

-- name: drop-table-artists
DROP TABLE IF EXISTS artists;

-- name: drop-table-genres
DROP TABLE IF EXISTS genres;

-- name: drop-table--artists-to-tracks
DROP TABLE IF EXISTS artists_to_tracks;

-- name: drop-table--artists-to-genres
DROP TABLE IF EXISTS artists_to_genres;

-- name: create-table-tracks
CREATE TABLE tracks
(
    id         VARCHAR(20) UNIQUE PRIMARY KEY NOT NULL,
    name       VARCHAR(255)                   NOT NULL,
    duration   INTEGER                        NOT NULL,
    explicit   INTEGER                        NOT NULL,
    popularity INTEGER                        NOT NULL
);

-- name: create-table-artists
CREATE TABLE artists
(
    id   VARCHAR(20) UNIQUE PRIMARY KEY NOT NULL,
    name VARCHAR(255)                   NOT NULL
);

-- name: create-table-genres
CREATE TABLE genres
(
    id   INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL
);

-- name: create-table-artists-to-tracks
CREATE TABLE artists_to_tracks
(
    id        INTEGER PRIMARY KEY AUTOINCREMENT,
    artist_id VARCHAR(20) NOT NULL,
    track_id  VARCHAR(20) NOT NULL
);

-- name: create-table-artists-to-genres
CREATE TABLE artists_to_genres
(
    id        INTEGER PRIMARY KEY AUTOINCREMENT,
    artist_id VARCHAR(20) NOT NULL,
    genre_id  INTEGER     NOT NULL

);

-- name: insert-track
INSERT OR
REPLACE
INTO tracks
VALUES (?, ?, ?, ?, ?);

-- name: insert-artist
INSERT OR
REPLACE
INTO artists
VALUES (?, ?);

-- name: insert-genre
INSERT OR
REPLACE
INTO genres
VALUES (?, ?);

-- name: insert-artist-to-track
INSERT OR
REPLACE
INTO artists_to_tracks
VALUES (?, ?, ?);

-- name: insert-artists-to-genres
INSERT OR
REPLACE
INTO artists_to_genres
VALUES (?, ?, ?);
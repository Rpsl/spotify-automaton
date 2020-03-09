package storage

import (
	"database/sql"

	"github.com/gchaincl/dotsql"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

type Storage struct {
	db  *sql.DB
	dot *dotsql.DotSql
}

const PathDatabase string = "db/base.db"
const PathDotsql string = "db/queries.sql"

func NewSqlite() (*Storage, error) {
	db, err := sql.Open("sqlite3", PathDatabase)

	if err != nil {
		return nil, err
	}

	dot, err := dotsql.LoadFromFile(PathDotsql)

	if err != nil {
		return nil, err
	}

	return &Storage{db: db, dot: dot}, nil
}

func (s *Storage) Init() error {
	migrations := []string{
		"drop-table-tracks",
		"drop-table-artists",
		"drop-table-genres",
		"drop-table--artists-to-tracks",
		"drop-table--artists-to-genres",
		"create-table-tracks",
		"create-table-artists",
		"create-table-genres",
		"create-table-artists-to-tracks",
		"create-table-artists-to-genres",
	}

	for _, migration := range migrations {
		log.Infof("exec migration: %s", migration)

		_, err := s.dot.Exec(s.db, migration)

		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Storage) AddTrack(track Track) error {
	_, err := s.dot.Exec(s.db, "insert-track", track.ID, track.Name, track.Duration, track.Explicit, track.Popularity)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) AddArtist(artist Artist) error {
	_, err := s.dot.Exec(s.db, "insert-artist", artist.ID, artist.Name)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetArtists() (*ArtistStack, error) {
	as := NewArtistStack()

	rows, err := s.dot.Query(s.db, "select-artists")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		a := Artist{}
		err := rows.Scan(&a.ID, &a.Name)

		if err != nil {
			continue
		}

		if rows.Err() != nil {
			continue
		}

		as.Add(a.ID, a.Name)
	}

	return as, nil
}

func (s *Storage) AddArtistToTrack(artistID string, trackID string) error {
	_, err := s.dot.Exec(s.db, "insert-artist-to-track", nil, artistID, trackID)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) AddGenre(genre string) (int64, error) {
	res, err := s.dot.Exec(s.db, "insert-genre", nil, genre)

	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (s *Storage) AddArtistToGenre(artistID string, genreID int64) error {
	_, err := s.dot.Exec(s.db, "insert-artists-to-genres", nil, artistID, genreID)

	if err != nil {
		return err
	}

	return nil
}

package cmd

import (
	log "github.com/sirupsen/logrus"

	"github.com/rpsl/spotify-automaton/config"
	"github.com/rpsl/spotify-automaton/spotify"
	"github.com/rpsl/spotify-automaton/storage"
)

type ContextApp struct {
	config  *config.Config
	storage *storage.Storage
	spotify *spotify.Client
}

func NewContext(config *config.Config) *ContextApp {
	db, err := storage.NewSqlite()

	if err != nil {
		log.Fatal(err)
	}

	client := spotify.New(config)

	return &ContextApp{
		config:  config,
		storage: db,
		spotify: client,
	}
}

func Refresh(config *config.Config) {
	context := NewContext(config)

	context.dbPrepare()
	context.dbRefresh()
}

func (c *ContextApp) dbPrepare() {
	err := c.storage.Init()

	if err != nil {
		log.Fatal(err)
	}
}

func (c *ContextApp) dbRefresh() {
	tracks, err := c.spotify.SavedTracks()

	if err != nil {
		log.Fatal(err)
	}

	ts := storage.NewTrackStack()
	as := storage.NewArtistStack()

	for _, track := range tracks {
		artists := make([]string, 0, len(track.Artists))

		for _, a := range track.Artists {
			as.Add(a.ID.String(), a.Name)
			artists = append(artists, a.ID.String())
		}

		ts.Add(track.ID.String(), track.Name, track.Duration, track.Explicit, track.Popularity, artists, track.AddedAt)
	}

	c.insertTracks(ts)
	c.insertArtists(as)
	c.updateGenres()
}

func (c *ContextApp) insertTracks(ts *storage.TrackStack) {
	for _, t := range ts.Tracks {
		err := c.storage.AddTrack(t)

		if err != nil {
			log.Errorf("error insert track %s :: %v", t.ID, err)
		}

		for _, a := range t.Artists {
			err := c.storage.AddArtistToTrack(a, t.ID)

			if err != nil {
				log.Errorf("error insert artist to track %s -> %s :: %v", t.ID, a, err)
			}
		}
	}
}

func (c *ContextApp) insertArtists(as *storage.ArtistStack) {
	for _, a := range as.Artists {
		err := c.storage.AddArtist(a)

		if err != nil {
			log.Errorf("error insert artist %s :: %v", a.ID, err)
		}
	}
}

func (c *ContextApp) updateGenres() {
	as, err := c.storage.GetArtists()

	if err != nil {
		log.Errorf("error select artist :: %v", err)
	}

	var (
		i      = 0
		chunk  []string
		chunks [][]string
	)

	for _, a := range as.Artists {
		chunk = append(chunk, a.ID)
		i++

		if len(chunk) == 50 {
			i = 0

			chunks = append(chunks, chunk)

			chunk = []string{}
		}
	}

	// last chunk
	chunks = append(chunks, chunk)

	for _, ids := range chunks {
		artists, err := c.spotify.Artists(ids)

		if err != nil {
			log.Errorf("error get full artist :: %v", err)
		}

		for _, artist := range artists {
			for _, genre := range artist.Genres {
				insID, _ := c.storage.AddGenre(genre)
				err = c.storage.AddArtistToGenre(artist.ID.String(), insID)

				if err != nil {
					log.Errorf("error add artist to genre :: %v", err)
				}
			}
		}
	}
}

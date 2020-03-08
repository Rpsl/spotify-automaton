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

		ts.Add(track.ID.String(), track.Name, track.Duration, track.Explicit, track.Popularity, artists)
	}

	c.insertTracks(ts)
	c.insertArtists(as)
}

func (c *ContextApp) insertTracks(ts *storage.TrackStack) {
	for _, t := range ts.Tracks {
		err := c.storage.AddTrack(t)

		if err != nil {
			log.Errorf("error insert track %s :: %v", t.ID, err)
		}

		for _, a := range t.Artists {
			err := c.storage.AddArtistToTrack(t.ID, a)

			if err != nil {
				log.Errorf("error insert artist to track %s -> %s :: %v", t.ID, a, err)
			}
		}
	}
}

// todo: holy shit. i need a sleep. an rewrite that
func (c *ContextApp) insertArtists(as *storage.ArtistStack) {
	for _, a := range as.Artists {
		err := c.storage.AddArtist(a)

		if err != nil {
			log.Errorf("error insert artist %s :: %v", a.ID, err)
		}
	}

	i := 0

	var ids []string

	for _, a := range as.Artists {
		i++

		ids = append(ids, a.ID)

		if i == 50 {
			res, err := c.spotify.Artists(ids)
			i = 0
			ids = nil

			if err != nil {
				log.Errorf("error get full artist :: %v", err)
			}

			for _, fa := range res {
				for _, fag := range fa.Genres {
					c.storage.AddGenre(fag)
					c.storage.AddArtistToGenre(fa.ID.String(), fag)
				}
			}
		}
	}
}

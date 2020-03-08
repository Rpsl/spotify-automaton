package spotify

import (
	"time"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"

	"github.com/rpsl/spotify-automaton/config"
)

type Client struct {
	sp spotify.Client
}

func New(config *config.Config) *Client {
	token := &oauth2.Token{
		AccessToken:  config.Tokens.AccessToken,
		RefreshToken: config.Tokens.RefreshToken,
		TokenType:    "Bearer",
		Expiry:       time.Date(2009, time.October, 7, 23, 0, 0, 0, time.UTC),
	}

	temp := spotify.NewAuthenticator("")
	temp.SetAuthInfo(config.Spotify.ClientID, config.Spotify.ClientSecret)
	client := temp.NewClient(token)
	client.AutoRetry = true

	return &Client{sp: client}
}

// todo spotify.* to TrackStack
func (c *Client) SavedTracks() ([]spotify.FullTrack, error) {
	var userTracks []spotify.FullTrack

	tracks, err := c.sp.CurrentUsersTracks()

	if err != nil {
		return userTracks, err
	}

	for {
		for _, track := range tracks.Tracks {
			userTracks = append(userTracks, track.FullTrack)
		}

		err = c.sp.NextPage(tracks)

		if err == spotify.ErrNoMorePages {
			return userTracks, nil
		}

		if err != nil {
			break
		}
	}

	return userTracks, err
}

func (c *Client) Artists(ids []string) ([]*spotify.FullArtist, error) {
	aID := make([]spotify.ID, 0, len(ids))

	for _, id := range ids {
		aID = append(aID, spotify.ID(id))
	}

	res, err := c.sp.GetArtists(aID...)

	return res, err
}

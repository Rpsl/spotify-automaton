package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

type Tasks struct {
	ID string `toml:"-"`
}

type Spotify struct {
	ClientID     string `toml:"spotify_client_id"`
	ClientSecret string `toml:"spotify_client_secret"`
	RedirectURL  string `toml:"spotify_redirect_url"`
}

type Tokens struct {
	AccessToken  string `toml:"access_token"`
	RefreshToken string `toml:"refresh_token"`
}

type Config struct {
	Spotify Spotify
	Tokens  Tokens
	Tasks   map[string]*Tasks
}

const PathConfig string = "config.toml"
const PathTokens string = "tokens.toml"

// LoadConfig loads TOML configuration from a file path
func LoadConfig() (*Config, error) {
	config := Config{}

	err := config.loadSpotify(PathConfig)

	if err != nil {
		log.Fatal(err)
	}

	err = config.loadTokens(PathTokens)

	if err != nil {
		log.Fatal(err)
	}

	return &config, nil
}

func (c *Config) WriteTokens(accessToken string, refreshToken string) {
	tokens, err := os.Create(PathTokens)

	if err != nil {
		log.Fatal(err)
	}

	defer tokens.Close()

	value := Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	coder := toml.NewEncoder(tokens)
	err = coder.Encode(value)

	if err != nil {
		log.Fatal(err)
	}
}

func (c *Config) loadSpotify(path string) error {
	_, err := toml.DecodeFile(path, &c.Spotify)

	if err != nil {
		return errors.Wrap(err, "failed to load spotify file")
	}

	return nil
}

func (c *Config) loadTokens(path string) error {
	var tokens Tokens

	_, err := toml.DecodeFile(path, &tokens)

	var pathError *os.PathError

	if err != nil {
		// ignore path error. we must be create toml file with login command
		if errors.As(err, &pathError) {
			return nil
		}

		return errors.Wrap(err, "failed to load tokens file")
	}

	c.Tokens = tokens

	return nil
}

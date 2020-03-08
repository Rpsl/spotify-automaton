package main

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/jawher/mow.cli"

	"github.com/rpsl/spotify-automaton/cmd"
	"github.com/rpsl/spotify-automaton/config"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: time.RFC3339,
		FullTimestamp:   true,
	})

	// Load TOML file
	// log.Debugf("loading configuration %q", config.PathConfig)
	cfg, err := config.LoadConfig()

	if err != nil {
		log.WithError(err).Fatal("failed to load configuration file")
	}

	app := cli.App("Spotify Automaton", "Utility for automation some tasks in your spotify account")

	app.Command("login", "get credential for login", func(c *cli.Cmd) {
		c.Action = func() {
			cmd.Login(cfg)
		}
	})

	app.Command("refresh", "refresh local database", func(c *cli.Cmd) {
		c.Action = func() {
			cmd.Refresh(cfg)
		}
	})

	app.Run(os.Args)
}

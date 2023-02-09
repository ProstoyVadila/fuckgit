package main

import (
	conf "github.com/ProstoyVadila/fuckgit/config"
	"github.com/rs/zerolog/log"
)

func init() {
	conf.SetLogger()
}

func main() {
	log.Info().Msg("Loading config...")
	config, err := conf.New("./")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	config.SetGinMode()

	log.Info().Msg("Starting server...")
	app, err := NewApp(config)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}

	app.Run()
}

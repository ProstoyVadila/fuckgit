package main

import (
	conf "github.com/ProstoyVadila/fuckgit/config"
	"github.com/rs/zerolog/log"
)

func init() {
	conf.SetLogger()
}

func main() {
	log.Info().Msg("Loading config")
	config, err := conf.New("./")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	config.SetGinMode()

	app, err := NewApp(config)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}

	log.Info().Str("address", config.Addr).Msg("Starting server")
	app.Run()
}

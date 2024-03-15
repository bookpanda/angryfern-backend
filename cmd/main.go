package main

import (
	"github.com/bookpanda/angryfern-backend/cfgldr"
	"github.com/bookpanda/angryfern-backend/database"

	"github.com/rs/zerolog/log"
)

func main() {
	conf, err := cfgldr.LoadConfig()
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", "file").
			Msg("Failed to load config")
	}

	_, err = database.New(&conf.DatabaseConfig)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", "file").
			Msg("Failed to init postgres connection")
	}
}

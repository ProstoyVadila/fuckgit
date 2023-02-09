package config

import (
	"time"

	"github.com/rs/zerolog"
)

func SetLogger(logLevels ...zerolog.Level) {
	var logLevel zerolog.Level
	if len(logLevels) == 0 {
		logLevel = zerolog.DebugLevel
	} else {
		logLevel = logLevels[0]
	}
	zerolog.TimeFieldFormat = zerolog.TimestampFunc().UTC().Format(time.RFC3339)
	zerolog.SetGlobalLevel(logLevel)
}

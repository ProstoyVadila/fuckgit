package config

import (
	"github.com/rs/zerolog"
)

const TimeFormat = "02.01.2006 15:04:05.000000000Z07"

func SetLogger(logLevels ...zerolog.Level) {
	var logLevel zerolog.Level
	if len(logLevels) == 0 {
		logLevel = zerolog.DebugLevel
	} else {
		logLevel = logLevels[0]
	}
	zerolog.TimeFieldFormat = TimeFormat
	zerolog.SetGlobalLevel(logLevel)
}

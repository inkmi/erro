package erro

import (
	"github.com/rs/zerolog"
)

var (
	LogTo *zerolog.Logger = nil

	DevMode = false

	//DefaultLogger logger implements default configuration for a logger
	DefaultLogger = &logger{
		config: &Config{
			LinesBefore: 4,
			LinesAfter:  2,
		},
	}
)

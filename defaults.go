package erro

import (
	"github.com/rs/zerolog"
)

var (
	LogTo *zerolog.Logger = nil

	DevMode = false

	DefaultConfig = Config{
		LinesBefore: 4,
		LinesAfter:  2,
	}
	//DefaultLogger logger implements default configuration for a logger
	DefaultLogger = &logger{
		config: &Config{
			LinesBefore: 4,
			LinesAfter:  2,
		},
	}
)

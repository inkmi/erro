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
)

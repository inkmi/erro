package internal

import (
	"github.com/rs/zerolog"
)

type printer func(format string, data ...interface{})

var (
	Printer printer         = nil
	LogTo   *zerolog.Logger = nil

	DevMode = false

	DefaultConfig = config{
		LinesBefore: 4,
		LinesAfter:  2,
	}
)

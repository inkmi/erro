package internal

import (
	"github.com/rs/zerolog"
)

// Config holds the configuration for a logger
type Config struct {
	LinesBefore int //How many lines to print *before* the error line when printing source code
	LinesAfter  int //How many lines to print *after* the error line when printing source code
}
type printer func(format string, data ...interface{})

var (
	Printer printer         = nil
	LogTo   *zerolog.Logger = nil

	DevMode = false

	DefaultConfig = Config{
		LinesBefore: 0,
		LinesAfter:  0,
	}
)

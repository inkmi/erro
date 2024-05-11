package main

import (
	"errors"

	"github.com/rs/zerolog/log"
)

func main() {
	logger := log.With().Caller().Logger()

	logger.Info().Int("Hello", 42).Msg("Info example")
	logger.Debug().Str("Hello", "🦄").Msg("Debug example")
	logger.Warn().Str("Hello", "World").Msg("Warn example")

	err := errors.New("Testerror")
	logger.Error().Err(err).Str("Test", "Test").Msg("Error example")
}

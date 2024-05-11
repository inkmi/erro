package main

import (
	"errors"

	"github.com/rs/zerolog/log"
)

func main() {
	logger := log.With().Caller().Logger()

	logger.Info().Msg("Welcome to erro 🧑‍🚀")
	logger.Info().Msg("https://github.com/inkmi/erro")
	logger.Info().Int("Hello", 42).Msg("Info example")
	logger.Trace().Int("Hello", 23).Msg("Trace example")
	logger.Debug().Str("Hello", "🦄").Msg("Debug example")
	logger.Warn().Str("Hello", "World").Msg("Warn example")

	err := errors.New("Testerror")
	logger.Error().Err(err).Str("Test", "Test").Msg("Error example")
	logger.Info().Int("After", 1).Int("Days", 2).Msg("After the error")
}

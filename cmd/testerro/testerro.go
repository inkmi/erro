package main

import (
	"errors"

	"github.com/rs/zerolog/log"
)

func main() {
	logger := log.With().Caller().Logger()

	err := errors.New("Testerror")
	logger.Error().Err(err).Str("Test", "Test").Msg("Test")
}

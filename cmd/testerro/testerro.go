package main

import (
	"errors"
	"fmt"

	"github.com/rs/zerolog/log"
)

func main() {
	logger := log.With().Caller().Logger()

	fmt.Println("Hello World")

	err := errors.New("Testerror")
	logger.Error().Err(err).Str("Test", "Test").Msg("Test")
}

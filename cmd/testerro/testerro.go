package main

import (
	"errors"
	"fmt"

	"github.com/rs/zerolog/log"
)

func main() {
	logger := log.With().Caller().Logger()
	log.Logger = logger

	logger.Info().Msg("Welcome to erro 🧑‍🚀")
	logger.Info().Msg("https://github.com/inkmi/erro")
	logger.Info().Int("Hello", 42).Msg("Info example")
	logger.Trace().Int("Hello", 23).Msg("Trace example")
	logger.Debug().Str("Hello", "🦄").Msg("Debug example")
	logger.Warn().Str("Hello", "World").Msg("Warn example")

	err := someBigFunction(2)
	logger.Error().Err(err).Str("Test", "Test").Msg("Error example")
	logger.Info().Int("After", 1).Int("Days", 2).Msg("After the error")
}

func someBigFunction(y int) error {
	x := 3
	someDumbFunction()

	someSmallFunction()

	someDumbFunction()

	x = 2
	if err := someNastyFunction(x, y); err != nil {
		log.Error().Err(err).Int("x", x).Int("y", y).Msg("Error example")
		return errors.New("can't call nasty function")
	}

	someSmallFunction()

	someDumbFunction()

	return errors.New("x")
}

func someSmallFunction() {
}

func someNastyFunction(x int, y int) error {
	return fmt.Errorf("i'm failing for no reason with %d and %d", x, y)
}

func someDumbFunction() bool {
	return false
}

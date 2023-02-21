package main

import (
	"errors"
	"fmt"
	"github.com/StephanSchmidt/erro"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	logger := log.With().Caller().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	erro.LogTo = &logger
	erro.DevMode = true

	wrapingFunc()

}

func wrapingFunc() {
	someBigFunction()
}

func someBigFunction() error {
	someDumbFunction()

	someSmallFunction()

	someDumbFunction()

	if e := someNastyFunction(); e != nil {
		return erro.New(e, "Can't open bug database")
	}

	someSmallFunction()

	someDumbFunction()

	return errors.New("x")
}

func someSmallFunction() {
	fmt.Print("I do things !")
}

func someNastyFunction() error {
	return errors.New("I'm failing for no reason")
}

func someDumbFunction() bool {
	return false
}

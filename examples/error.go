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
	someBigFunction(2)
}

func someBigFunction(y int) error {
	someDumbFunction()

	someSmallFunction()

	someDumbFunction()

	x := 3
	if e := someNastyFunction(x, y); e != nil {
		return erro.Errorf(e, "Can't call nasty function")
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

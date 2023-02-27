package main

import (
	"errors"
	"fmt"
	"github.com/StephanSchmidt/erro"
	"github.com/StephanSchmidt/erro/internal"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	logger := log.With().Caller().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	internal.LogTo = &logger
	internal.DevMode = true

	wrapingFunc()

}

func wrapingFunc() {
	someBigFunction(2)
}

func someBigFunction(y int) error {
	x := 3
	someDumbFunction()

	someSmallFunction()

	someDumbFunction()

	x = 2
	if e := someNastyFunction(x, y); e != nil {
		return erro.Errorf("Can't call nasty function", e, x)
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

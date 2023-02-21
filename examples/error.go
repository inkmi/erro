package main

import (
	"erro"
	"errors"
	"github.com/sirupsen/logrus"
)

var (
	debug = erro.NewLogger(&erro.Config{
		PrintFunc:               logrus.Errorf,
		LinesBefore:             6,
		LinesAfter:              3,
		PrintError:              true,
		PrintSource:             true,
		DisableStackIndentation: false,
		PrintStack:              true,
		ExitOnDebugSuccess:      true,
	})
)

func main() {
	erro.DevMode = true

	logrus.Print("Start of the program")

	wrapingFunc()

	logrus.Print("End of the program")
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
	logrus.Print("I do things !")
}

func someNastyFunction() error {
	return errors.New("I'm failing for no reason")
}

func someDumbFunction() bool {
	return false
}

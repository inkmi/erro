// Package erro provides a simple object to enhance Go source code debugging
//
// Example result:
//
//	$ go run myfailingapp.go
//	Program starting
//	error in main.main: something failed here
//	line 13 of /Users/StephanSchmidt/go/src/github.com/StephanSchmidt/sandbox/testerr.go
//	9: func main() {
//	10:     fmt.Println("Program starting")
//	11:     err := errors.Errorf("something failed here")
//	12:
//	13:     erro.Debug(err)
//	14:
//	15:     fmt.Println("End of the program")
//	16: }
//	exit status 1
//
// You can configure your own logger with these options :
//
//	type Config struct {
//		LinesBefore        int
//		LinesAfter         int
//		PrintStack         bool
//		getData        bool
//		PrintError         bool
//		ExitOnDebugSuccess bool
//	}
//
// Example :
//
//	debug := erro.NewLogger(&erro.Config{
//		LinesBefore:        2,
//		LinesAfter:         1,
//		PrintError:         true,
//		getData:        true,
//		PrintStack:         false,
//		ExitOnDebugSuccess: true,
//	})
//
//	// ...
//	if err != nil {
//		debug.Debug(err)
//		return
//	}
//
// Outputs :
//
//	Error in main.someBigFunction(): I'm failing for no reason
//	line 41 of /Users/StephanSchmidt/go/src/github.com/StephanSchmidt/sandbox/testerr.go:41
//	33: func someBigFunction() {
//	...
//	40:     if err := someNastyFunction(); err != nil {
//	41:             debug.Debug(err)
//	42:             return
//	43:     }
//	exit status 1
package erro

import (
	"errors"
	"fmt"
	"github.com/StephanSchmidt/erro/internal"
	"github.com/rs/zerolog"
)

func SetDevModeOn() {
	internal.DevMode = true
}

func SetDevModeOff() {
	internal.DevMode = false
}

func SetLogger(logger *zerolog.Logger) {
	internal.LogTo = logger
}

func Errorf(format string, source error, a ...any) error {
	err := internal.PrintErro(source, a...)
	if err != nil {
		return err
	}
	return fmt.Errorf(format, source, a)
}

func New(errorString string, source error, a ...interface{}) error {
	err := internal.PrintErro(source, a...)
	if err != nil {
		return err
	}
	n := errors.New(errorString)
	return errors.Join(n, source)
}

func NewE(myError error, source error, a ...interface{}) error {
	err := internal.PrintErro(source, a...)
	if err != nil {
		return err
	}
	return errors.Join(myError, source)
}

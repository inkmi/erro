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
//		printSource        bool
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
//		printSource:        true,
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
	"github.com/spf13/afero"
)

var (
	debugMode = false
	fs        = afero.NewOsFs() //fs is at package level because I think it needn't be scoped to loggers
)

func Errorf(format string, source error, a ...interface{}) error {
	DefaultLogger.Overload(1) // Prevents from adding this func to the stack trace
	return DefaultLogger.Errorf(format, source, a...)
}

func New(errorString string, source error, a ...interface{}) error {
	DefaultLogger.Overload(1) // Prevents from adding this func to the stack trace
	return DefaultLogger.New(errorString, source, a...)
}

func NewE(myError error, source error, a ...interface{}) error {
	DefaultLogger.Overload(1) // Prevents from adding this func to the stack trace
	return DefaultLogger.NewE(myError, source, a...)
}

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
	"fmt"

	"github.com/StephanSchmidt/erro/internal"
	"github.com/rs/zerolog"
)

type ErroError struct {
	err  error
	data map[string]any
}

type ErrorData interface {
	ErrorData() map[string]any
}

func (ee ErroError) Error() string {
	return ""
}

func (ee ErroError) ErrorData() map[string]any {
	return ee.data
}

func (ee *ErroError) Str(key string, value any) *ErroError {
	ee.data[key] = value
	return ee
}

func (ee ErroError) LogTo(log zerolog.Logger) {
	LogError(log, ee)
}

func LogError(log zerolog.Logger, err ErroError) {
	event := log.Error()
	for k, v := range err.ErrorData() {
		event = event.Interface(k, v)
	}
	event.Err(err)
}

func SetDevModeOn() {
	internal.DevMode = true
}

func SetDevModeOff() {
	internal.DevMode = false
}

func SetLogger(logger *zerolog.Logger) {
	internal.LogTo = logger
}

// ErrorF("User %s could not be found", err, userId).Str("UserId", userId)

func Errorf(format string, a ...any) ErroError {
	if len(a) > 0 {
		// Attempt to type assert the first argument to 'error'.
		if errObject, ok := a[0].(error); ok {
			err := internal.PrintErro(errObject, a[1:]...)
			if err != nil {
				// Do nothing
			}
			ee := ErroError{
				err:  fmt.Errorf(format, errObject, a),
				data: make(map[string]any),
			}
			return ee
		}
	}
	ee := ErroError{
		err:  fmt.Errorf(format, a...),
		data: make(map[string]any),
	}
	return ee
}

func Wrap(source error, a ...interface{}) ErroError {
	e := internal.PrintErro(source, a...)
	if e != nil {
		// Do nothing
	}
	ee := ErroError{
		err:  source,
		data: make(map[string]any),
	}
	return ee
}

func New(err error, source error, a ...interface{}) ErroError {
	e := internal.PrintErro(source, a...)
	if e != nil {
		// Do nothing
	}
	ee := ErroError{
		err:  err,
		data: make(map[string]any),
	}
	return ee
}

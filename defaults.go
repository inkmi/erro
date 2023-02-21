package erro

import (
	"fmt"
	"github.com/rs/zerolog"
)

var (
	//DefaultLoggerPrintFunc is fmt.Printf without return values
	DefaultLoggerPrintFunc = func(format string, data ...interface{}) {
		fmt.Printf(format+"\n", data...)
	}

	LogTo *zerolog.Logger = nil

	DevMode = false
	//DefaultLogger logger implements default configuration for a logger
	DefaultLogger = &logger{
		config: &Config{
			PrintFunc:          DefaultLoggerPrintFunc,
			LinesBefore:        4,
			LinesAfter:         2,
			PrintStack:         true,
			PrintSource:        true,
			PrintError:         true,
			ExitOnDebugSuccess: false,
		},
	}
)

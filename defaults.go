package erro

import (
	"fmt"
	"github.com/rs/zerolog"
)

var (
	//DefaultLoggerPrintFunc is fmt.printf without return values
	DefaultLoggerPrintFunc = func(format string, data ...interface{}) {
		fmt.Printf(format+"\n", data...)
	}

	LogTo *zerolog.Logger = nil

	DevMode = false

	//DefaultLogger logger implements default configuration for a logger
	DefaultLogger = &logger{
		config: &Config{
			LinesBefore: 4,
			LinesAfter:  2,
		},
	}
)

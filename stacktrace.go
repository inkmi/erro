package erro

import (
	"runtime/debug"
	"strconv"
	"strings"
)

// StackTraceItem represents parsed information of a stack trace item
type StackTraceItem struct {
	CallingObject string
	Args          []string
	SourcePathRef string
	SourceLineRef int
	MysteryNumber int64 // don't know what this is, no documentation found, if you know please let me know via a PR !
}

func parseStackTrace(deltaDepth int) []StackTraceItem {
	return parseAnyStackTrace(string(debug.Stack()), deltaDepth)
}

func parseAnyStackTrace(stackStr string, deltaDepth int) []StackTraceItem {
	stackArr := strings.Split(stackStr, "\n")
	if len(stackArr) < 2*(2+deltaDepth) {
		return nil
	}
	stack := strings.Join(stackArr[2*(2+deltaDepth):], "\n") //get stack trace and reduce to desired size
	parsedRes := regexpParseStack.FindAllStringSubmatch(stack, -1)

	sti := make([]StackTraceItem, len(parsedRes))
	for i := range parsedRes {
		args := regexpHexNumber.FindAllString(parsedRes[i][2], -1)
		srcLine, err := strconv.Atoi(parsedRes[i][4])
		if err != nil {
			srcLine = -1
		}

		mysteryNumberStr := parsedRes[i][5]
		mysteryNumber := int64(-25)
		if mysteryNumberStr != "" {
			mysteryNumber, err = strconv.ParseInt(parsedRes[i][5], 16, 32)
			if err != nil {
				mysteryNumber = -1
			}
		}

		sti[i] = StackTraceItem{
			CallingObject: parsedRes[i][1],
			Args:          args,
			SourcePathRef: parsedRes[i][3],
			SourceLineRef: srcLine,
			MysteryNumber: mysteryNumber,
		}
	}

	// get rid of the stacktrace item with our framework code
	return sti[1:]
}

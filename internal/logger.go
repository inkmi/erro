package internal

import (
	"errors"
	"github.com/fatih/color"
	"runtime/debug"
	"strings"
)

// config holds the configuration for a logger
type config struct {
	LinesBefore int //How many lines to print *before* the error line when printing source code
	LinesAfter  int //How many lines to print *after* the error line when printing source code
}

func PrintErro(source error, a ...any) error {
	if DevMode {
		if source == nil {
			return errors.New("erro: no error given")
		}

		stack := string(debug.Stack())
		stackItems := parseAnyStackTrace(stack, 0)
		if stackItems == nil || len(stackItems) < 1 {
			printf("Error: %s", source)
			printf("Erro tried to debug the error but the stack trace seems empty. If you think this is an error, please open an issue at https://github.com/StephanSchmidt/erro/issues/new and provide us logs to investigate.")
			return errors.New("erro can't find a stack")
		}
		stackSourceItemIndex := 0
		fileName := stackItems[stackSourceItemIndex].SourcePathRef
		callingObject := stackItems[stackSourceItemIndex].CallingObject
		debugLine := stackItems[stackSourceItemIndex].SourceLineRef

		// Print Source code
		lines := ReadSource(fileName)
		if lines == nil || len(lines) == 0 {
			return errors.New("erro can't read source")
		}

		data := getData(lines, fileName, debugLine, a, DefaultConfig.LinesBefore, DefaultConfig.LinesAfter)
		data.Stack = stackItems

		printf("Error in %s: %s", callingObject, color.YellowString(source.Error()))

		if data.FailingLine != -1 {
			printf("line %d of %s:%d", data.FailingLine+1, data.ShortFileName, data.FailingLine+1)
		} else {
			printf("error in %s (failing line not found, stack trace says func call is at line %d)", data.ShortFileName, data.DebugLine)
		}
		printSource(lines, data)
		printUsedVariables(data.UsedVars)
		printStack(data.Stack)
	}
	return nil
}

// DebugSource prints certain lines of source code of a file for debugging, using (*logger).config as configurations
func getData(lines []string, file string, debugLineNumber int,
	varValues []interface{},
	linesBefore int,
	linesAfter int,
) PrintSourceOptions {
	//find func line and adjust minLine if below
	funcLine := findFuncLine(lines, debugLineNumber)
	failingLineIndex, columnStart, columnEnd := findFailingLine(lines, funcLine, debugLineNumber)

	funcSrc := strings.Join(lines[funcLine:FindEndOfFunction(lines, funcLine)+1], "\n")

	var argNames []string
	if debugLineNumber > -1 {
		argNames = findArgNames(lines[debugLineNumber-1])[2:]
	}
	var failingArgs []string
	if failingLineIndex > -1 {
		failingArgs = extractArgs(lines[failingLineIndex][columnStart:])
	}
	usedVars := findUsedArgsLastWrite(funcLine, funcSrc, lines, argNames, varValues, failingArgs)

	data := PrintSourceOptions{
		DebugLine:     debugLineNumber,
		ShortFileName: getShortFilePath(file),
		FailingLine:   failingLineIndex,
		FuncLine:      funcLine,
		Highlighted: map[int][]int{
			failingLineIndex: {columnStart, columnEnd},
		},
		StartLine: max(funcLine, debugLineNumber-linesBefore),
		EndLine:   debugLineNumber + linesAfter,
		UsedVars:  usedVars,
	}
	return data
}

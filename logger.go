package erro

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"runtime/debug"
	"strings"
)

// Logger interface allows to log an error, or to print source code lines. Check out NewLogger function to learn more about Logger objects and Config.
type Logger interface {
	Errorf(fmt string, err error, a ...interface{}) error
	New(errorString string, err error, a ...interface{}) error
	NewE(myErr error, source error, a ...interface{}) error
}

// Config holds the configuration for a logger
type Config struct {
	LinesBefore int //How many lines to print *before* the error line when printing source code
	LinesAfter  int //How many lines to print *after* the error line when printing source code
}

// logger holds logger object, implementing Logger interface
type logger struct {
	config *Config //config for the logger
}

func (l *logger) New(errorString string, source error, a ...interface{}) error {
	err := printErro(l, source, a)
	if err != nil {
		return err
	}
	n := errors.New(errorString)
	return errors.Join(n, source)
}

func (l *logger) NewE(myErr error, source error, a ...interface{}) error {
	err := printErro(l, source, a)
	if err != nil {
		return err
	}
	return errors.Join(myErr, source)
}

func (l *logger) Errorf(format string, source error, a ...any) error {
	err := printErro(l, source, a)
	if err != nil {
		return err
	}
	return fmt.Errorf(format, source, a)
}

func printErro(l *logger, source error, a []any) error {
	if DevMode {
		if source == nil {
			return errors.New("erro: no error given")
		}

		stack := string(debug.Stack())
		stackItems := parseAnyStackTrace(stack, 1)
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

		data := l.getData(lines, fileName, debugLine, a)
		data.Stack = stackItems

		printf("Error in %s: %s", callingObject, color.YellowString(source.Error()))
		l.printSource(lines, data)
		// Print Stack Trace
	}
	return nil
}

func (l *logger) printSource(lines []string, data PrintSourceOptions) {
	printSource(lines, data)
	printUsedVariables(data.UsedVars)
	printStack(data.Stack)
}

// DebugSource prints certain lines of source code of a file for debugging, using (*logger).config as configurations
func (l *logger) getData(lines []string, file string, debugLineNumber int, varValues []interface{}) PrintSourceOptions {
	//find func line and adjust minLine if below
	funcLine := findFuncLine(lines, debugLineNumber)
	failingLineIndex, columnStart, columnEnd := findFailingLine(lines, funcLine, debugLineNumber)

	if failingLineIndex != -1 {
		printf("line %d of %s:%d", failingLineIndex+1, GetShortFilePath(file), failingLineIndex+1)
	} else {
		printf("error in %s (failing line not found, stack trace says func call is at line %d)", GetShortFilePath(file), debugLineNumber)
	}

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
		FailingLine: failingLineIndex,
		FuncLine:    funcLine,
		Highlighted: map[int][]int{
			failingLineIndex: {columnStart, columnEnd},
		},
		StartLine: max(funcLine, debugLineNumber-l.config.LinesBefore),
		EndLine:   debugLineNumber + l.config.LinesAfter,
		UsedVars:  usedVars,
	}
	return data
}

// printf is the function used to log
func printf(format string, data ...interface{}) {
	if LogTo != nil {
		(*LogTo).Debug().Msgf(format, data...)
	}
}

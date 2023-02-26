package erro

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
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
		stLines := parseStackTrace(2)
		if stLines == nil || len(stLines) < 1 {
			printf("Error: %s", source)
			printf("Erro tried to debug the error but the stack trace seems empty. If you think this is an error, please open an issue at https://github.com/StephanSchmidt/erro/issues/new and provide us logs to investigate.")
			return errors.New("erro can't find a stack")
		}
		// print Error
		printf("Error in %s: %s", stLines[0].CallingObject, color.YellowString(source.Error()))

		// Print Source code
		lines := ReadSource(stLines[0].SourcePathRef)
		if lines == nil || len(lines) == 0 {
			return errors.New("erro can't read source")
		}
		data := l.getData(lines, stLines[0].SourcePathRef, stLines[0].SourceLineRef, a)
		l.printSource(lines, data)
		// Print Stack Trace
		printStack(stLines)
	}
	return nil
}

func (l *logger) printSource(lines []string, data PrintSourceOptions) {
	printSource(lines, data)
	printUsedVariables(data)
}

// DebugSource prints certain lines of source code of a file for debugging, using (*logger).config as configurations
func (l *logger) getData(lines []string, file string, debugLineNumber int, varValues []interface{}) PrintSourceOptions {
	//find func line and adjust minLine if below
	funcLine := findFuncLine(lines, debugLineNumber)
	failingLineIndex, columnStart, columnEnd, argNames := findFailingLine(lines, funcLine, debugLineNumber)

	if failingLineIndex != -1 {
		printf("line %d of %s:%d", failingLineIndex+1, GetShortFilePath(file), failingLineIndex+1)
	} else {
		printf("error in %s (failing line not found, stack trace says func call is at line %d)", GetShortFilePath(file), debugLineNumber)
	}

	funcSrc := strings.Join(lines[funcLine:FindEndOfFunction(lines, funcLine)+1], "\n")
	usedVars := getUsedVars(funcSrc, lines, funcLine, failingLineIndex, columnStart, argNames, varValues)

	data := PrintSourceOptions{
		FuncLine: funcLine,
		Highlighted: map[int][]int{
			failingLineIndex: {columnStart, columnEnd},
		},
		StartLine: max(funcLine, debugLineNumber-l.config.LinesBefore),
		EndLine:   debugLineNumber + l.config.LinesAfter,
		UsedVars:  usedVars,
	}
	return data
}

func findUsedArgsLastWrite(
	funcLine int,
	funcSrc string,
	src []string,
	argNames []string,
	varValues []interface{},
	failingArgs []string) []UsedVar {

	var usedVars []UsedVar
	for i, ar := range argNames {
		lastWrite := lastWriteToVar(funcSrc, ar)
		uv := UsedVar{
			Name:            ar,
			Value:           varValues[i],
			LastWrite:       lastWrite,
			SourceLastWrite: strings.TrimSpace(src[lastWrite+funcLine-1]),
		}
		usedVars = append(usedVars, uv)
	}
	for _, fa := range diff(failingArgs, argNames) {
		lastWrite := lastWriteToVar(funcSrc, fa)
		lastWriteSrc := ""
		if lastWrite > -1 {
			lastWrite = lastWrite + funcLine
			lastWriteSrc = strings.TrimSpace(src[lastWrite-1])
		}
		uv := UsedVar{
			Name:            fa,
			Value:           nil,
			LastWrite:       lastWrite,
			SourceLastWrite: lastWriteSrc,
		}
		usedVars = append(usedVars, uv)
	}
	return usedVars
}

// printf is the function used to log
func printf(format string, data ...interface{}) {
	if LogTo != nil {
		(*LogTo).Debug().Msgf(format, data...)
	}
}

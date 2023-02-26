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
	//PrintSource prints lines based on given opts (see PrintSourceOptions type definition)
	PrintSource(lines []string, opts PrintSourceOptions)
	//DebugSource debugs a source file
	DebugSource(lines []string, file string, debugLineNumber int, varValues []interface{})
}

// Config holds the configuration for a logger
type Config struct {
	LinesBefore int //How many lines to print *before* the error line when printing source code
	LinesAfter  int //How many lines to print *after* the error line when printing source code
}

// PrintSourceOptions represents config for (*logger).PrintSource func
type PrintSourceOptions struct {
	FuncLine    int
	StartLine   int
	EndLine     int
	Highlighted map[int][]int //map[lineIndex][columnstart, columnEnd] of chars to highlight
}

// logger holds logger object, implementing Logger interface
type logger struct {
	config             *Config //config for the logger
	stackDepthOverload int     //stack depth to ignore when reading stack
}

// NewLogger creates a new logger struct with given config
func NewLogger(cfg *Config) Logger {
	l := logger{
		config:             cfg,
		stackDepthOverload: 0,
	}
	return &l
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
		stLines := parseStackTrace(1 + l.stackDepthOverload)
		if stLines == nil || len(stLines) < 1 {
			l.Printf("Error: %s", source)
			l.Printf("Erro tried to debug the error but the stack trace seems empty. If you think this is an error, please open an issue at https://github.com/StephanSchmidt/erro/issues/new and provide us logs to investigate.")
			return errors.New("erro can't find a stack")
		}
		// print Error
		l.Printf("Error in %s: %s", stLines[0].CallingObject, color.YellowString(source.Error()))

		// Print Source code
		lines := ReadSource(stLines[0].SourcePathRef)
		if lines == nil || len(lines) == 0 {
			return errors.New("erro can't read source")
		}
		l.DebugSource(lines, stLines[0].SourcePathRef, stLines[0].SourceLineRef, a)

		// Print Stack Trace
		l.Printf(color.BlueString("Stack trace:"))
		printStack(stLines)

		l.stackDepthOverload = 0
	}
	return nil
}

// DebugSource prints certain lines of source code of a file for debugging, using (*logger).config as configurations
func (l *logger) DebugSource(lines []string, file string, debugLineNumber int, varValues []interface{}) {
	//find func line and adjust minLine if below
	funcLine := findFuncLine(lines, debugLineNumber)
	failingLineIndex, columnStart, columnEnd, argNames := findFailingLine(lines, funcLine, debugLineNumber)

	if failingLineIndex != -1 {
		l.Printf("line %d of %s:%d", failingLineIndex+1, GetShortFilePath(file), failingLineIndex+1)
	} else {
		l.Printf("error in %s (failing line not found, stack trace says func call is at line %d)", GetShortFilePath(file), debugLineNumber)
	}

	l.PrintSource(lines, PrintSourceOptions{
		FuncLine: funcLine,
		Highlighted: map[int][]int{
			failingLineIndex: {columnStart, columnEnd},
		},
		StartLine: max(funcLine, debugLineNumber-l.config.LinesBefore),
		EndLine:   debugLineNumber + l.config.LinesAfter,
	})

	// Use AST instead of strings and []string in the future
	funcSrc := strings.Join(lines[funcLine:FindEndOfFunction(lines, funcLine)+1], "\n")
	var failingArgs []string
	if failingLineIndex > -1 {
		failingArgs = extractArgs(lines[failingLineIndex][columnStart:])
	}
	printVariables(l, findUsedArgs(funcLine, funcSrc, lines, argNames, varValues, failingArgs))
}

func findUsedArgs(
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

// PrintSource prints source code based on opts
func (l *logger) PrintSource(lines []string, opts PrintSourceOptions) {
	PrintSource(lines, opts, l)
}

// Printf is the function used to log
func (l *logger) Printf(format string, data ...interface{}) {
	if LogTo != nil {
		(*LogTo).Debug().Msgf(format, data...)
	}
}

// Overload adds depths to remove when parsing next stack trace
func (l *logger) Overload(amount int) {
	l.stackDepthOverload += amount
}

package erro

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/afero"
	"os"
	"path/filepath"
	"strings"
)

var (
	gopath = os.Getenv("GOPATH")
)

// Logger interface allows to log an error, or to print source code lines. Check out NewLogger function to learn more about Logger objects and Config.
type Logger interface {
	Errorf(fmt string, err error, a ...interface{}) error
	New(errorString string, err error, a ...interface{}) error
	//PrintSource prints lines based on given opts (see PrintSourceOptions type definition)
	PrintSource(lines []string, opts PrintSourceOptions)
	//DebugSource debugs a source file
	DebugSource(filename string, lineNumber int, args []any)
	//SetConfig replaces current config with the given one
	SetConfig(cfg *Config)
	//Config returns current config
	Config() *Config
	//Disable is used to disable Logger (every call to this Logger will perform NO-OP (no operation)) and return instantly
	//Use Disable(true) to disable and Disable(false) to enable again
	Disable(bool)
}

// Config holds the configuration for a logger
type Config struct {
	PrintFunc               func(format string, data ...interface{}) //Printer func (eg: fmt.Printf)
	LinesBefore             int                                      //How many lines to print *before* the error line when printing source code
	LinesAfter              int                                      //How many lines to print *after* the error line when printing source code
	PrintStack              bool                                     //Shall we print stack trace ? yes/no
	PrintSource             bool                                     //Shall we print source code along ? yes/no
	PrintError              bool                                     //Shall we print the error of Debug(err) ? yes/no
	ExitOnDebugSuccess      bool                                     //Shall we os.Exit(1) after Debug has finished logging everything ? (doesn't happen when err is nil)
	DisableStackIndentation bool                                     //Shall we print stack vertically instead of indented
	Mode                    int
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

	l.Doctor()

	return &l
}

func (l *logger) New(errorString string, source error, a ...interface{}) error {
	if DevMode {
		l.Doctor()
		if source == nil {
			return errors.New("Erro: no error given")
		}

		stLines := parseStackTrace(1 + l.stackDepthOverload)
		if stLines == nil || len(stLines) < 1 {
			l.Printf("Error: %s", source)
			l.Printf("Erro tried to debug the error but the stack trace seems empty. If you think this is an error, please open an issue at https://github.com/StephanSchmidt/erro/issues/new and provide us logs to investigate.")
			return errors.New("Erro: can't find a stack")
		}

		// print Error
		l.Printf("Error in %s: %s", stLines[0].CallingObject, color.YellowString(source.Error()))
		// print Source lines
		l.DebugSource(stLines[0].SourcePathRef, stLines[0].SourceLineRef, a)

		// print Stack trace
		l.Printf(color.BlueString("Stack trace:"))
		l.printStack(stLines)

		l.stackDepthOverload = 0
	}
	n := errors.New(errorString)
	return errors.Join(n, source)
}

func (l *logger) NewE(myErr error, source error, a ...interface{}) error {
	if DevMode {
		l.Doctor()
		if myErr == nil || source == nil {
			return errors.New("Erro: no error or no source given")
		}

		stLines := parseStackTrace(1 + l.stackDepthOverload)
		if stLines == nil || len(stLines) < 1 {
			l.Printf("Error: %s", myErr)
			l.Printf("Erro tried to debug the error but the stack trace seems empty. If you think this is an error, please open an issue at https://github.com/StephanSchmidt/erro/issues/new and provide us logs to investigate.")
			return errors.New("Erro: can't find a stack")
		}

		// print Error
		l.Printf("Error in %s: %s", stLines[0].CallingObject, color.YellowString(source.Error()))
		// print Source lines
		l.DebugSource(stLines[0].SourcePathRef, stLines[0].SourceLineRef, a)
		// print Stack trace
		l.Printf(color.BlueString("Stack trace:"))
		l.printStack(stLines)

		l.stackDepthOverload = 0
	}
	return errors.Join(myErr, source)
}

func (l *logger) Errorf(format string, source error, a ...any) error {
	if DevMode {
		l.Doctor()
		if source == nil {
			return errors.New("Erro: no error given")
		}

		stLines := parseStackTrace(1 + l.stackDepthOverload)
		if stLines == nil || len(stLines) < 1 {
			l.Printf("Error: %s", source)
			l.Printf("Erro tried to debug the error but the stack trace seems empty. If you think this is an error, please open an issue at https://github.com/StephanSchmidt/erro/issues/new and provide us logs to investigate.")
			return errors.New("Erro can't find a stack")
		}

		// print Error
		l.Printf("Error in %s: %s", stLines[0].CallingObject, color.YellowString(source.Error()))

		// Print Source code
		l.DebugSource(stLines[0].SourcePathRef, stLines[0].SourceLineRef, a)

		// Print Stack Trace
		l.Printf(color.BlueString("Stack trace:"))
		l.printStack(stLines)

		l.stackDepthOverload = 0
	}
	return fmt.Errorf(format, source, a)
}

// DebugSource prints certain lines of source code of a file for debugging, using (*logger).config as configurations
func (l *logger) DebugSource(filepath string, debugLineNumber int, args []interface{}) {
	filepathShort := filepath
	if gopath != "" {
		filepathShort = strings.Replace(filepath, gopath+"/src/", "", -1)
	}

	b, err := afero.ReadFile(fs, filepath)
	if err != nil {
		l.Printf("erro: cannot read file '%s': %s. If sources are not reachable in this environment, you should set PrintSource=false in logger config.", filepath, err)
		return
		// l.Debug(err)
	}
	lines := strings.Split(string(b), "\n")

	// set line range to print based on config values and debugLineNumber
	minLine := debugLineNumber - l.config.LinesBefore
	maxLine := debugLineNumber + l.config.LinesAfter

	//delete blank lines from range and clean range if out of lines range
	deleteBlankLinesFromRange(lines, &minLine, &maxLine)

	//free some memory from unused values
	lines = lines[:maxLine+1]

	//find func line and adjust minLine if below
	funcLine := findFuncLine(lines, debugLineNumber)
	if funcLine > minLine {
		minLine = funcLine + 1
	}

	//try to find failing line if any
	failingLineIndex, columnStart, columnEnd, argNames := findFailingLine(lines, funcLine, debugLineNumber)
	failingArgs := extractArgs(lines[failingLineIndex][columnStart:])

	if failingLineIndex != -1 {
		l.Printf("line %d of %s:%d", failingLineIndex+1, filepathShort, failingLineIndex+1)
	} else {
		l.Printf("error in %s (failing line not found, stack trace says func call is at line %d)", filepathShort, debugLineNumber)
	}

	l.PrintSource(lines, PrintSourceOptions{
		FuncLine: funcLine,
		Highlighted: map[int][]int{
			failingLineIndex: {columnStart, columnEnd},
		},
		StartLine: minLine,
		EndLine:   maxLine,
	})

	if len(argNames) > 0 {
		l.Printf(color.BlueString("Variables:"))
		for i, arg := range argNames {
			l.Printf("  %v : %v", arg, args[i])
		}
		// Print all args that were in the failing call but not in the Errorf/New call
		for _, arg := range diff(failingArgs, argNames) {
			l.Printf("  %v : ?", arg)
		}
	}
}

func diff(first []string, second []string) []string {
	var diff []string
	for _, str := range first {
		if !contains(second, str) {
			diff = append(diff, str)
		}
	}
	return diff
}

func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

// PrintSource prints source code based on opts
func (l *logger) PrintSource(lines []string, opts PrintSourceOptions) {
	//print func on first line
	if opts.FuncLine != -1 && opts.FuncLine < opts.StartLine {
		l.Printf("%s", color.RedString("%d: %s", opts.FuncLine+1, lines[opts.FuncLine]))
		if opts.FuncLine < opts.StartLine-1 { // append blank line if minLine is not next line
			l.Printf("%s", color.YellowString("..."))
		}
	}

	for i := opts.StartLine; i < opts.EndLine; i++ {
		if _, ok := opts.Highlighted[i]; !ok || len(opts.Highlighted[i]) != 2 {
			l.Printf("%d: %s", i+1, color.YellowString(lines[i]))
			continue
		}

		hlStart := max(opts.Highlighted[i][0], 0)          //highlight column start
		hlEnd := min(opts.Highlighted[i][1], len(lines)-1) //highlight column end
		l.Printf("%d: %s%s%s", i+1, color.YellowString(lines[i][:hlStart]), color.RedString(lines[i][hlStart:hlEnd+1]), color.YellowString(lines[i][hlEnd+1:]))
	}
}

func (l *logger) Doctor() (neededDoctor bool) {
	neededDoctor = false

	if l.config.PrintFunc == nil {
		neededDoctor = true
		if LogTo != nil {
			(*LogTo).Debug().Msg("PrintFunc not set for this logger. Replacing with DefaultLoggerPrintFunc.")
		}
		l.config.PrintFunc = DefaultLoggerPrintFunc
	}

	if l.config.LinesBefore < 0 {
		neededDoctor = true
		if LogTo != nil {
			(*LogTo).Debug().Msgf("LinesBefore is '%d' but should not be <0. Setting to 0.", l.config.LinesBefore)
		}
		l.config.LinesBefore = 0
	}

	if l.config.LinesAfter < 0 {
		neededDoctor = true
		if LogTo != nil {
			(*LogTo).Debug().Msgf("LinesAfters is '%d' but should not be <0. Setting to 0.", l.config.LinesAfter)
		}
		l.config.LinesAfter = 0
	}

	if neededDoctor && !debugMode && LogTo != nil {
		(*LogTo).Debug().Msgf("erro: Doctor() has detected and fixed some problems on your logger configuration. It might have modified your configuration. Check logs by enabling debug. 'erro.SetDebugMode(true)'.")
	}
	return
}

func (l *logger) printStack(stLines []StackTraceItem) {
	for i := len(stLines) - 1; i >= 0; i-- {
		padding := ""
		if !l.config.DisableStackIndentation {
			for j := 0; j < len(stLines)-1-i-1; j++ {
				padding += "  "
			}
			if i < len(stLines)-1 {
				padding += "╰╴"
			}
		}
		if LogTo != nil {
			file := filepath.Base(stLines[i].SourcePathRef)
			(*LogTo).Debug().Msgf(padding+"%s ( %s:%d )", stLines[i].CallingObject, file, stLines[i].SourceLineRef)
		}
	}
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

func (l *logger) SetConfig(cfg *Config) {
	l.config = cfg
	l.Doctor()
}

func (l *logger) Config() *Config {
	return l.config
}

func (l *logger) SetMode(mode int) bool {
	if !isIntInSlice(mode, enabledModes) {
		return false
	}
	l.Config().Mode = mode
	return true
}

func (l *logger) Disable(shouldDisable bool) {
	if shouldDisable {
		l.Config().Mode = ModeDisabled
	} else {
		l.Config().Mode = ModeEnabled
	}
}

const (
	// ModeDisabled represents the disabled mode (NO-OP)
	ModeDisabled = iota + 1
	// ModeEnabled represents the enabled mode (Print)
	ModeEnabled
)

var (
	enabledModes = []int{ModeDisabled, ModeEnabled}
)

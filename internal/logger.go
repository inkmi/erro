package internal

import (
	"errors"
	"fmt"
	"runtime/debug"
	"strings"

	"github.com/fatih/color"
)

func FindErrorOrigin(lines []string, logLine int) ([]int, error) {
	var errVarName string

	// Parse the log line to find the error variable name
	logContents := lines[logLine-1] // Adjust for 0-based indexing in slices
	parts := strings.Split(logContents, ".")
	for _, part := range parts {
		if strings.Contains(part, "Err(") {
			errVarName = strings.Trim(strings.Split(part, "Err(")[1], ") ")
			break
		}
	}

	if errVarName == "" {
		return nil, fmt.Errorf("error variable not found in the log statement")
	}

	// Scan backwards to find where this variable was defined or assigned
	for i := logLine - 2; i >= 0; i-- {
		line := strings.ReplaceAll(lines[i], " ", "") // Removing spaces to simplify detection
		if strings.Contains(line, errVarName) {
			// Check if this line includes an assignment that impacts the error variable
			if strings.Contains(line, ":=") || strings.Contains(line, "=") {
				assignmentPart := strings.SplitAfter(line, "=")[0]
				if strings.Contains(assignmentPart, ":=") {
					assignmentPart = strings.SplitAfter(assignmentPart, ":=")[0]
				}

				// Parse all variables in the assignment
				vars := strings.Split(strings.Split(assignmentPart, "=")[0], ",")
				for _, v := range vars {
					if strings.Contains(v, errVarName) {
						// Now find the complete statement, handling multiline statements
						startLine := i
						for startLine > 0 && strings.TrimSpace(lines[startLine-1]) == "" {
							startLine--
						}
						endLine := i + 1
						for endLine < len(lines) && strings.TrimSpace(lines[endLine]) == "" {
							endLine++
						}
						return []int{startLine, endLine - 1}, nil // Convert to 1-based index for line numbers
					}
				}
			}
		}
	}

	return nil, fmt.Errorf("no assignment for %s found before the log statement", errVarName)
}

func ExtractFilename(fullPath string) string {
	// Split the path by the "/" character
	parts := strings.Split(fullPath, "/")

	// The last element will be the filename with line number
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}

	// If there's no "/" or input is empty, return the original input
	return fullPath
}

func PrintError(fileName string, debugLine int) error {
	// Print Source code
	lines := ReadSource(fileName)
	if len(lines) == 0 {
		return errors.New("erro can't read source")
	}

	red := color.New(color.FgHiRed).SprintFunc()

	assignment, err := FindErrorOrigin(lines, debugLine)
	if err != nil {
		return err
	}
	if len(assignment) == 0 {
		fmt.Printf("%s%s%s%s %s\n", red(ExtractFilename(fileName)), red(":"), red(debugLine), red(":"), red(strings.TrimSpace(lines[debugLine-1])))

	} else {
		lineNo := assignment[0] + 1
		fmt.Printf("%s%s%s%s %s\n", red(ExtractFilename(fileName)), red(":"), red(lineNo+1), red(":"), red(strings.TrimSpace(lines[lineNo])))
	}

	return nil
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
		if len(lines) == 0 {
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

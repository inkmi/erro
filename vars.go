package erro

import (
	"github.com/fatih/color"
)

type UsedVar struct {
	Name            string
	Value           interface{}
	LastWrite       int
	SourceLastWrite string
}

func getUsedVars(funcSrc string, lines []string, funcLine int, failingLineIndex int, columnStart int, argNames []string, varValues []interface{}) []UsedVar {
	// Use AST instead of strings and []string in the future
	var failingArgs []string
	if failingLineIndex > -1 {
		failingArgs = extractArgs(lines[failingLineIndex][columnStart:])
	}
	return findUsedArgsLastWrite(funcLine, funcSrc, lines, argNames, varValues, failingArgs)
}

func printUsedVariables(data PrintSourceOptions) {
	if len(data.UsedVars) > 0 {
		printf(color.BlueString("Variables:"))
		for _, arg := range data.UsedVars {
			if arg.Value != nil {
				printf(" %v : %v", arg.Name, arg.Value)
			} else {
				printf(" %v : ?", arg.Name)
			}
			if arg.LastWrite > -1 {
				printf(" ╰╴ %d : %v", arg.LastWrite, arg.SourceLastWrite)
			}
		}
	}
}

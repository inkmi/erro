package internal

import (
	"strings"

	"github.com/fatih/color"
)

type UsedVar struct {
	Name            string
	Value           interface{}
	LastWrite       int
	SourceLastWrite string
}

func GoFindUsedArgsLastWrite(
	funcLine int,
	funcSrc string,
	src []string,
	failingArgs []string) []UsedVar {

	var usedVars []UsedVar
	for _, fa := range failingArgs {
		if isValidVariable(fa) {
			lastWrite := GoLastWriteToVar(funcSrc, fa)
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
	}
	return usedVars
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
		lastWrite := GoLastWriteToVar(funcSrc, ar)
		uv := UsedVar{
			Name:            ar,
			Value:           varValues[i],
			LastWrite:       lastWrite,
			SourceLastWrite: strings.TrimSpace(src[lastWrite+funcLine-1]),
		}
		usedVars = append(usedVars, uv)
	}
	for _, fa := range diff(failingArgs, argNames) {
		lastWrite := GoLastWriteToVar(funcSrc, fa)
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

func printUsedVariables(vars []UsedVar) {
	if len(vars) > 0 {
		printf(color.BlueString("Variables:"))
		for _, arg := range vars {
			if arg.Value != nil {
				printf(" %v : %v\n", arg.Name, arg.Value)
			} else {
				printf(" %v : ?\n", arg.Name)
			}
			if arg.LastWrite > -1 {
				printf(" ╰╴ %d : %v\n", arg.LastWrite, arg.SourceLastWrite)
			}
		}
	}
}

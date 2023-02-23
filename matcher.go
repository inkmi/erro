package erro

import (
	"regexp"
	"strings"
)

// Debug regex with https://regex101.com/

var (
	regexpFuncLine                   = regexp.MustCompile(`^func\s[a-zA-Z0-9]+\(([^)]*)\)\s+[a-zA-Z0-9]*\s*{`)
	regexpParseDebugLineParseVarName = regexp.MustCompile(`erro\.Errorf\(|erro\.New\(|erro\.NewE\(`)
)

func MatchFunc(line string) bool {
	return regexpFuncLine.Match([]byte(line))
}

func ArgNames(line string) []string {
	reMatches := regexpParseDebugLineParseVarName.FindIndex([]byte(line))
	matches := len(reMatches)
	if matches > 0 {
		args := extractArgs(line[matches:])
		for i, _ := range args {
			args[i] = strings.TrimSpace(args[i])
		}
		return args
	} else {
		return nil
	}
}

func MatchVarName(line string) *string {
	reMatches := regexpParseDebugLineParseVarName.FindIndex([]byte(line))
	matches := len(reMatches)
	if matches > 0 {
		args := extractArgs(line[matches:])
		errorName := strings.TrimSpace(args[1])
		return &errorName
	} else {
		return nil
	}
}

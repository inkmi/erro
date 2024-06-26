package internal

import (
	"regexp"
	"strings"
)

// Debug regex with https://regex101.com/

var (
	regexpFuncLine                   = regexp.MustCompile(`^func\s[a-zA-Z0-9]+\(([^)]*)\)\s+[a-zA-Z0-9]*[^{]*{`)
	regexpParseDebugLineParseVarName = regexp.MustCompile(`erro\.Errorf\(|erro\.New\(|erro\.Wrap\(`)
)

func MatchFunc(line string) bool {
	return regexpFuncLine.Match([]byte(line))
}

func findArgNames(line string) []string {
	reMatches := regexpParseDebugLineParseVarName.FindIndex([]byte(line))
	matches := len(reMatches)
	if matches > 0 {
		args := extractArgs(line[matches:])
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

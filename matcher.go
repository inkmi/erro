package erro

import (
	"regexp"
	"strings"
)

// Debug regex with https://regex101.com/

var (
	regexpFuncLine                   = regexp.MustCompile(`^func\s[a-zA-Z0-9]+\(([^)]*)\)\s+[a-zA-Z0-9]*\s*{`)
	regexpParseDebugLineParseVarName = regexp.MustCompile(`erro\.Errorf\(([^,]*),([^,)]*)|erro\.New\(([^,]*),([^,)]*)|erro\.NewE\(([^,]*),([^,)]*)`)
)

func MatchFunc(line string) bool {
	return regexpFuncLine.Match([]byte(line))
}

func MatchVarName(line string) *string {
	reMatches := regexpParseDebugLineParseVarName.FindStringSubmatch(line)
	matches := len(reMatches)
	if matches < 2 {
		return nil
	} else {
		if matches > 2 && reMatches[2] != "" {
			varName := strings.TrimSpace(reMatches[2])
			return &varName
		} else if matches > 4 && reMatches[4] != "" {
			varName := strings.TrimSpace(reMatches[4])
			return &varName
		} else if matches > 6 && reMatches[6] != "" {
			varName := strings.TrimSpace(reMatches[6])
			return &varName
		} else {
			return nil
		}
	}
}

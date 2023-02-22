package erro

import (
	"regexp"
)

// Debug regex with https://regex101.com/

var (
	regexpFuncLine                   = regexp.MustCompile(`^func\s[a-zA-Z0-9]+\(([^)]*)\)\s+[a-zA-Z0-9]*\s*{`)
	regexpParseDebugLineParseVarName = regexp.MustCompile(`erro\.Errorf\(([^,]*)|erro\.New\(([^,]*)`)
)

func MatchFunc(line string) bool {
	return regexpFuncLine.Match([]byte(line))
}

func MatchVarName(line string) *string {
	reMatches := regexpParseDebugLineParseVarName.FindStringSubmatch(line)
	if len(reMatches) < 2 {
		return nil
	} else {
		if reMatches[1] == "" {
			return &reMatches[2]
		} else {
			return &reMatches[1]
		}
	}
}

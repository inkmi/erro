package erro

import "strings"

func extractArgs(s string) []string {
	var args []string
	start := strings.Index(s, "(")
	end := strings.LastIndex(s, ")")
	if start != -1 && end != -1 && end > start {
		argStr := s[start+1 : end]
		args = strings.Split(argStr, ",")
	}
	return args
}

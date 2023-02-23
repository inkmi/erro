package erro

import "strings"

func splitWithBraces(s string, sep rune) []string {
	var result []string
	var stack []rune
	var start int
	for i, c := range s {
		switch c {
		case '(':
			stack = append(stack, c)
		case ')':
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				panic("unmatched closing brace")
			}
		case sep:
			if len(stack) == 0 {
				result = append(result, s[start:i])
				start = i + 1
			}
		}
	}
	if len(stack) != 0 {
		panic("unmatched opening brace")
	}
	result = append(result, s[start:])
	return result
}

func extractArgs(s string) []string {
	var args []string
	start := strings.Index(s, "(")
	end := strings.LastIndex(s, ")")
	if start != -1 && end != -1 && end > start {
		argStr := s[start+1 : end]
		args = splitWithBraces(argStr, ',')
	}
	for i, _ := range args {
		args[i] = strings.TrimSpace(args[i])
	}
	return args
}

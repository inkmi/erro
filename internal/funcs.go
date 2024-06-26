package internal

import (
	"strings"
)

// findFuncLine finds line where func is declared
func findFuncLine(lines []string, lineNumber int) int {
	for i := lineNumber; i >= 0; i-- {
		if MatchFunc(lines[i]) {
			return i
		}
	}
	return -1
}

func OpeningClosePos(s string) (int, int) {
	var stack []rune
	start := -1
	end := -1
	for i, c := range s {
		switch c {
		case '(':
			if start == -1 && len(stack) == 0 {
				start = i
			}
			stack = append(stack, c)
		case ')':
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				panic("unmatched closing brace")
			}
			if end == -1 && len(stack) == 0 {
				end = i
			}
		}
	}
	return start, end
}

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
				result = append(result, strings.TrimSpace(s[start:i]))
				start = i + 1
			}
		}
	}
	if len(stack) != 0 {
		panic("unmatched opening brace")
	}
	result = append(result, strings.TrimSpace(s[start:]))
	return result
}

func extractArgs(s string) []string {
	var args []string
	start, end := OpeningClosePos(s)
	if start != -1 && end != -1 && end > start {
		argStr := s[start+1 : end]
		args = splitWithBraces(argStr, ',')
	}
	for i := range args {
		args[i] = strings.TrimSpace(args[i])
	}
	return args
}

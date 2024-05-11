package internal

import (
	"fmt"
	"strings"
)

func GolangFindErrorOrigin(lines []string, logLine int) ([]int, error) {
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
	for i := logLine - 1; i >= 0; i-- {
		line := strings.ReplaceAll(lines[i], " ", "") // Removing spaces to simplify detection
		if strings.Contains(line, errVarName) {
			// Check if this line includes an assignment that impacts the error variable
			if strings.Contains(line, ":=") || strings.Contains(line, "=") {
				assignmentPart := strings.SplitAfter(line, "=")[0]
				if strings.Contains(assignmentPart, "!=") {
					continue
				}
				if strings.Contains(assignmentPart, ":=") {
					assignmentPart = strings.SplitAfter(assignmentPart, ":=")[0]
				}
				// Parse all variables in the assignment
				vars := strings.Split(strings.Split(assignmentPart, "=")[0], ",")
				for _, v := range vars {
					if strings.Contains(v, errVarName) {
						// Now find the complete statement, handling multiline statements
						startLine := i
						endLine := i
						return []int{startLine, endLine - 1}, nil // Convert to 1-based index for line numbers
					}
				}
			}
		}
	}

	return nil, fmt.Errorf("no assignment for %s found before the log statement", errVarName)
}

package internal

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

func isValidVariable(s string) bool {
	// Regex for matching valid variable names.
	// It starts with a letter or underscore and can be followed by any number of letters, digits, or underscores.
	var validVariablePattern = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)
	return validVariablePattern.MatchString(s)
}

func PrintError(fileName string, debugLine int) error {
	nomralizedFileName := strings.ToLower(strings.TrimSpace(ExtractFilename(fileName)))

	supported := true
	if !strings.HasSuffix(nomralizedFileName, ".go") {
		supported = false
	}

	lines := ReadSource(fileName)
	if len(lines) == 0 {
		return errors.New("erro can't read source")
	}

	if supported {
		data, err := GoGetData(lines, fileName, debugLine, DefaultConfig.LinesBefore, DefaultConfig.LinesAfter)

		if err != nil {
			return err
		}
		red := color.New(color.FgHiRed).SprintFunc()

		for _, uv := range data.UsedVars {
			if isValidVariable(uv.Name) {
				fmt.Printf("    %s%s%s %s%s %s\n", red(ExtractFilename(fileName)), red(":"), red(uv.LastWrite),
					red(uv.Name),
					red(":"),
					red(uv.SourceLastWrite),
				)
			}
		}

		if data.FailingLine == 0 {
			fmt.Printf("    %s%s%s %s\n", red(ExtractFilename(fileName)), red(":"), red(debugLine), red(strings.TrimSpace(lines[debugLine-1])))

		} else {
			fmt.Printf("    %s%s%s %s\n", red(ExtractFilename(fileName)), red(":"), red(data.FailingLine+1), red(strings.TrimSpace(lines[data.FailingLine])))
		}

	} else {
		red := color.New(color.FgHiRed).SprintFunc()
		fmt.Printf("    %s%s%s%s %s\n", red(ExtractFilename(fileName)), red(":"), red(debugLine), red(":"), red(strings.TrimSpace(lines[debugLine-1])))
	}

	return nil
}

package internal

import (
	"errors"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

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
		assignment, err := GolangFindErrorOrigin(lines, debugLine)
		if err != nil {
			return err
		}
		red := color.New(color.FgHiRed).SprintFunc()
		if len(assignment) == 0 {
			fmt.Printf("%s%s%s%s %s\n", red(ExtractFilename(fileName)), red(":"), red(debugLine), red(":"), red(strings.TrimSpace(lines[debugLine-1])))

		} else {
			lineNo := assignment[0] + 1
			fmt.Printf("%s%s%s%s %s\n", red(ExtractFilename(fileName)), red(":"), red(lineNo+1), red(":"), red(strings.TrimSpace(lines[lineNo])))
		}
	} else {
		red := color.New(color.FgHiRed).SprintFunc()
		fmt.Printf("%s%s%s%s %s\n", red(ExtractFilename(fileName)), red(":"), red(debugLine), red(":"), red(strings.TrimSpace(lines[debugLine-1])))
	}

	return nil
}

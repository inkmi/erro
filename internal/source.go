package internal

import (
	"sort"
	"strings"

	"github.com/fatih/color"
)

func printSource(lines []string, opts PrintSourceOptions) {
	var filteredVars []UsedVar
	for _, v := range opts.UsedVars {
		if !strings.HasPrefix(v.SourceLastWrite, "func") {
			filteredVars = append(filteredVars, v)
		}
	}

	sort.Slice(filteredVars, func(i, j int) bool {
		return filteredVars[i].LastWrite < filteredVars[j].LastWrite
	})

	// First print the surrounding function
	if opts.FuncLine != -1 && opts.FuncLine < opts.StartLine {
		printf("    %s:%d %s\n", opts.ShortFileName, opts.FuncLine+1, color.RedString("%s", strings.TrimSpace(lines[opts.FuncLine])))
		if opts.FuncLine < opts.StartLine-1 { // append blank line if minLine is not next line
			printf("    %s\n", color.YellowString("..."))
		}
	}

	//if len(filteredVars) > 0 {
	//	printf("    %s:%d %s\n", opts.ShortFileName, filteredVars[0].LastWrite, color.RedString(filteredVars[0].SourceLastWrite))
	//
	//	// Iterate over the slice starting from the second item
	//	for i := 1; i < len(filteredVars); i++ {
	//		// Check if there is a gap between the current and the previous item
	//		if filteredVars[i].LastWrite > filteredVars[i-1].LastWrite+1 {
	//			println(color.YellowString("..."))
	//		}
	//		printf("    %s:%d %s\n", opts.ShortFileName, filteredVars[i].LastWrite, color.RedString(filteredVars[i].SourceLastWrite))
	//	}
	//
	//	if filteredVars[len(filteredVars)-1].LastWrite < opts.FailingLine {
	//		printf("    %s\n", color.YellowString("..."))
	//	}
	//}

	// Last print the failing line
	printf("    %s:%d %s\n", opts.ShortFileName, opts.FailingLine+1, color.RedString(strings.TrimSpace(lines[opts.FailingLine])))

}

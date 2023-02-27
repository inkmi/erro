package internal

import "github.com/fatih/color"

func printSource(lines []string, opts PrintSourceOptions) {
	//print func on first line
	if opts.FuncLine != -1 && opts.FuncLine < opts.StartLine {
		printf("%s", color.RedString("%d: %s", opts.FuncLine+1, lines[opts.FuncLine]))
		if opts.FuncLine < opts.StartLine-1 { // append blank line if minLine is not next line
			printf("%s", color.YellowString("..."))
		}
	}

	for i := opts.StartLine; i < opts.EndLine; i++ {
		if _, ok := opts.Highlighted[i]; !ok || len(opts.Highlighted[i]) != 2 {
			printf("%d: %s", i+1, color.YellowString(lines[i]))
			continue
		}

		hlStart := max(opts.Highlighted[i][0], 0)          //highlight column start
		hlEnd := min(opts.Highlighted[i][1], len(lines)-1) //highlight column end
		printf("%d: %s%s%s", i+1, color.YellowString(lines[i][:hlStart]), color.RedString(lines[i][hlStart:hlEnd+1]), color.YellowString(lines[i][hlEnd+1:]))
	}
}

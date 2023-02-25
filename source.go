package erro

import "github.com/fatih/color"

func PrintSource(lines []string, opts PrintSourceOptions, l *logger) {
	//print func on first line
	if opts.FuncLine != -1 && opts.FuncLine < opts.StartLine {
		l.Printf("%s", color.RedString("%d: %s", opts.FuncLine+1, lines[opts.FuncLine]))
		if opts.FuncLine < opts.StartLine-1 { // append blank line if minLine is not next line
			l.Printf("%s", color.YellowString("..."))
		}
	}

	for i := opts.StartLine; i < opts.EndLine; i++ {
		if _, ok := opts.Highlighted[i]; !ok || len(opts.Highlighted[i]) != 2 {
			l.Printf("%d: %s", i+1, color.YellowString(lines[i]))
			continue
		}

		hlStart := max(opts.Highlighted[i][0], 0)          //highlight column start
		hlEnd := min(opts.Highlighted[i][1], len(lines)-1) //highlight column end
		l.Printf("%d: %s%s%s", i+1, color.YellowString(lines[i][:hlStart]), color.RedString(lines[i][hlStart:hlEnd+1]), color.YellowString(lines[i][hlEnd+1:]))
	}
}

package erro

import "github.com/fatih/color"

type UsedVar struct {
	Name            string
	Value           interface{}
	LastWrite       int
	SourceLastWrite string
}

func PrintVariables(l *logger, vars []UsedVar) {
	if len(vars) > 0 {
		l.Printf(color.BlueString("Variables:"))
		for _, arg := range vars {
			if arg.Value != nil {
				l.Printf(" %v : %v", arg.Name, arg.Value)
			} else {
				l.Printf(" %v : ?", arg.Name)
			}
			if arg.LastWrite > -1 {
				l.Printf(" ╰╴ %d : %v", arg.LastWrite, arg.SourceLastWrite)
			}
		}
	}
}

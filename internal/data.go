package internal

// PrintSourceOptions represents config for (*logger).getData func
type PrintSourceOptions struct {
	ShortFileName string
	FuncLine      int
	FailingLine   int
	LogLine       int
	StartLine     int
	EndLine       int
	Highlighted   map[int][]int
	UsedVars      []UsedVar
	Stack         []StackTraceItem
}

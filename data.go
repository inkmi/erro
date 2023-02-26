package erro

// PrintSourceOptions represents config for (*logger).getData func
type PrintSourceOptions struct {
	FuncLine    int
	FailingLine int
	StartLine   int
	EndLine     int
	Highlighted map[int][]int
	UsedVars    []UsedVar
	Stack       []StackTraceItem
}

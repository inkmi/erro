package erro

// PrintSourceOptions represents config for (*logger).getData func
type PrintSourceOptions struct {
	FuncLine    int
	StartLine   int
	EndLine     int
	Highlighted map[int][]int
	UsedVars    []UsedVar
	//map[lineIndex][columnstart, columnEnd] of chars to highlight
}

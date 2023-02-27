package internal

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// printf is the function used to log
func printf(format string, data ...interface{}) {
	if LogTo != nil {
		(*LogTo).Debug().Msgf(format, data...)
	}
}

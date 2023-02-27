package internal

import (
	"fmt"
)

type TestPrinter struct {
	Output []string
}

func NewTestPrinter() *TestPrinter {
	tp := TestPrinter{
		Output: make([]string, 0),
	}
	return &tp
}

func (t *TestPrinter) Printf(format string, data ...any) {
	output := fmt.Sprintf(format, data...)
	t.Output = append(t.Output, output)
}

func TestPrinterFunc(t *TestPrinter) printer {
	return func(format string, data ...any) {
		t.Printf(format, data...)
	}
}

package internal

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintSource(t *testing.T) {
	tp := NewTestPrinter()
	Printer = TestPrinterFunc(tp)
	lines := strings.Split(`func test() {
x = 4
y = 3
err := add(x,y)
Errorf("x", err, x, y)
}
`, "\n")
	data := PrintSourceOptions{
		LogLine:       4,
		ShortFileName: getShortFilePath("error.go"),
		FailingLine:   3,
		FuncLine:      0,
		Highlighted:   map[int][]int{},
		StartLine:     0,
		EndLine:       len(lines),
		UsedVars:      []UsedVar{},
	}
	printSource(lines, data)
	assert.Equal(t, 1, len(tp.Output))
}

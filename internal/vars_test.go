package internal

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsedArgs(t *testing.T) {
	t.Parallel()
	code := `func add(x int, y int) int {
	x := 2
    x = 4
    err := calc(x,y)
    erro.New("x", err, x)
}
`
	erroArgs := []string{"x"}
	failingArgs := []string{"x", "y"}
	vals := []any{2}
	u := findUsedArgsLastWrite(0, code, strings.Split(code, "\n"), erroArgs, vals, failingArgs)
	u1 := UsedVar{
		Name:            "x",
		Value:           2,
		LastWrite:       3,
		SourceLastWrite: "x = 4",
	}
	u2 := UsedVar{
		Name:            "y",
		Value:           nil,
		LastWrite:       1,
		SourceLastWrite: "func add(x int, y int) int {",
	}
	assert.Equal(t, 2, len(u))
	assert.Equal(t, u1, u[0])
	assert.Equal(t, u2, u[1])
}

func TestPrintUsedVariables(t *testing.T) {
	tp := NewTestPrinter()
	u1 := UsedVar{
		Name:            "x",
		Value:           2,
		LastWrite:       10,
		SourceLastWrite: "x = 2",
	}
	u2 := UsedVar{
		Name:            "y",
		Value:           nil,
		LastWrite:       15,
		SourceLastWrite: "y = 2",
	}
	Printer = TestPrinterFunc(tp)
	printUsedVariables([]UsedVar{u1, u2})
	assert.Equal(t, 5, len(tp.Output))
	assert.Equal(t, "Variables:", tp.Output[0])
	assert.Equal(t, " x : 2", tp.Output[1])
	assert.Equal(t, " ╰╴ 10 : x = 2", tp.Output[2])
	assert.Equal(t, " y : ?", tp.Output[3])
	assert.Equal(t, " ╰╴ 15 : y = 2", tp.Output[4])

}

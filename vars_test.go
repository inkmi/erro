package erro

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
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
	args := []string{"x"}
	failingArgs := []string{"y"}
	vals := []any{2}
	u := findUsedArgsLastWrite(0, code, strings.Split(code, "\n"), args, vals, failingArgs)
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

package internal

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGetData(t *testing.T) {
	t.Parallel()
	code := strings.Split(
		`func add(x int, y int) int {
	x := 2
    x = 4
    err := calc(x,y)
    if err != nil {
		return erro.Errorf("x", err, x)
    }
    return x
}`, "\n")
	data := getData(
		code,
		"/stephan/src/example.go",
		6,
		[]any{4},
		4,
		2,
	)
	assert.Equal(t, 3, data.FailingLine)
	assert.Equal(t, 0, data.FuncLine)
	assert.Equal(t, 2, data.StartLine)
	assert.Equal(t, 8, data.EndLine)
	assert.Equal(t, 6, data.DebugLine)
	uv := []UsedVar{
		{
			"x",
			4,
			3,
			"x = 4",
		},
		{
			"y",
			nil,
			1,
			"func add(x int, y int) int {",
		},
	}
	assert.Equal(t, uv, data.UsedVars)
}

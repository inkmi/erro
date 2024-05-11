package internal

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetData(t *testing.T) {
	t.Parallel()
	code := strings.Split(
		`func add(x int, y int) (int,error) {
	x := 2
    x = 4
    err := calc(x,y)
    if err != nil {
		return 0, erro.Errorf("x", err, x).Str("UserId", 1)
    }
    return x, nil
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
			"func add(x int, y int) (int,error) {",
		},
	}
	assert.Equal(t, uv, data.UsedVars)
}

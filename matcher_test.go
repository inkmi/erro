package erro

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFuncNoArgsMatch(t *testing.T) {
	t.Parallel()
	line := "func a() {   "
	assert.Equal(t, true, MatchFunc(line))
}

func TestFuncNoArgsReturnMatch(t *testing.T) {
	t.Parallel()
	line := "func a() string {   "
	assert.Equal(t, true, MatchFunc(line))
}

func TestFuncOneArgMatch(t *testing.T) {
	t.Parallel()
	line := "func a(x int) {   "
	assert.Equal(t, true, MatchFunc(line))
}

func TestVarName(t *testing.T) {
	t.Parallel()
	line := "erro.Errorf(\"Can't call nasty function\", e)\n"
	varName := MatchVarName(line)
	assert.NotNil(t, varName)
	if varName != nil {
		assert.Equal(t, "e", *varName)
	}
}

func TestVarNameNew(t *testing.T) {
	t.Parallel()
	line := "erro.New(\"Can't call nasty function\", e)\n"
	varName := MatchVarName(line)
	assert.NotNil(t, varName)
	if varName != nil {
		assert.Equal(t, "e", *varName)
	}
}

func TestVarNameNewE(t *testing.T) {
	t.Parallel()
	line := "erro.NewE(e1, e2)\n"
	varName := MatchVarName(line)
	assert.NotNil(t, varName)
	if varName != nil {
		assert.Equal(t, "e2", *varName)
	}
}

func TestFuncArgNames(t *testing.T) {
	t.Parallel()
	line := `erro.Errorf("abc", err, a, b, c){   `
	argNames := ArgNames(line)
	assert.Equal(t, 5, len(argNames))
	assert.Equal(t, "err", argNames[1])
	assert.Equal(t, "a", argNames[2])
	assert.Equal(t, "b", argNames[3])
	assert.Equal(t, "c", argNames[4])
}

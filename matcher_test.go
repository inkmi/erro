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

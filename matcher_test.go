package erro

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFuncNoArgsMatch(t *testing.T) {
	line := "func a() {   "
	assert.Equal(t, true, MatchFunc(line))
}

func TestFuncNoArgsReturnMatch(t *testing.T) {
	line := "func a() string {   "
	assert.Equal(t, true, MatchFunc(line))
}

func TestFuncOneArgMatch(t *testing.T) {
	line := "func a(x int) {   "
	assert.Equal(t, true, MatchFunc(line))
}

func TestVarName(t *testing.T) {
	line := "erro.Errorf(e, \"Can't call nasty function\")\n"
	varName := MatchVarName(line)
	assert.NotNil(t, varName)
	if varName != nil {
		assert.Equal(t, "e", *varName)
	}
}

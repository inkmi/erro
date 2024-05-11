package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestFuncArgNames(t *testing.T) {
	t.Parallel()
	line := `erro.Errorf("abc", err, a, b, c){   `
	argNames := findArgNames(line)
	assert.Equal(t, 5, len(argNames))
	assert.Equal(t, "err", argNames[1])
	assert.Equal(t, "a", argNames[2])
	assert.Equal(t, "b", argNames[3])
	assert.Equal(t, "c", argNames[4])
}

func TestExtractArgs(t *testing.T) {
	t.Parallel()
	line := `x := call("abc", err, a, b, c)`
	argNames := extractArgs(line)
	assert.Equal(t, 5, len(argNames))
	assert.Equal(t, "err", argNames[1])
	assert.Equal(t, "a", argNames[2])
	assert.Equal(t, "b", argNames[3])
	assert.Equal(t, "c", argNames[4])
}

func TestExtractArgsWithFunctionCall(t *testing.T) {
	t.Parallel()
	line := `x := call("abc", err, add(2,3), b, c)`
	argNames := extractArgs(line)
	assert.Equal(t, 5, len(argNames))
	assert.Equal(t, "err", argNames[1])
	assert.Equal(t, "add(2,3)", argNames[2])
	assert.Equal(t, "b", argNames[3])
	assert.Equal(t, "c", argNames[4])
}

func TestSplitWithBraces(t *testing.T) {
	t.Parallel()
	line := `add(2,3), b, c`
	res := splitWithBraces(line, ',')

	assert.Equal(t, 3, len(res))
	assert.Equal(t, "add(2,3)", res[0])
	assert.Equal(t, "b", res[1])
	assert.Equal(t, "c", res[2])
}

func TestFindFuncLine(t *testing.T) {
	lines := []string{
		"x := 2",
		"func add(x int) (x, error) {",
		"  if x < 0 { return errors.Errorf(\"Error\") }" +
			"  return x + 1, nil",
		"}",
	}
	funcLine := findFuncLine(lines, 1)
	assert.Equal(t, 1, funcLine)
}

func TestFindOpeningClose(t *testing.T) {
	start, end := OpeningClosePos("    (a,b,c,(x,y),d).(f.g)")
	assert.Equal(t, 4, start)
	assert.Equal(t, 18, end)
}

func TestFindFuncLineBeforeFuncReturnsError(t *testing.T) {
	lines := []string{
		"x := 2",
		"func add(x int) {",
		"  return x + 1",
		"}",
	}
	funcLine := findFuncLine(lines, 0)
	assert.Equal(t, -1, funcLine)
}

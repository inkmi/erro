package erro

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestLastWrite(t *testing.T) {
	src := `
func add(y int) int {
   x := 3
   return y + x
}`
	lastWrite := lastWriteToVar(src, "x")
	assert.Equal(t, 3, lastWrite)
}

func TestLastWriteNotExists(t *testing.T) {
	src := `
func add(y int) int {
   x := 3
   return y + x
}`
	lastWrite := lastWriteToVar(src, "z")
	assert.Equal(t, -1, lastWrite)
}

func TestLastWriteTwoWrites(t *testing.T) {
	src := `
func add(y int) int {
   x := 2
   x = 3
   return y + x
}`
	lastWrite := lastWriteToVar(src, "x")
	assert.Equal(t, 4, lastWrite)
}

func TestLastWriteFunc(t *testing.T) {
	src := `
func add(y int) int {
   x := 2
   x = 3
   return y + x
}`
	lastWrite := lastWriteToVar(src, "y")
	assert.Equal(t, 2, lastWrite)
}

func TestFindEndOfFunction(t *testing.T) {
	src := `
func add(y int) int {
   x := 3
   z := { 
      3
   }
   return y + x
}

func sub(y int) int {
   x := 3
   return y - x
}
`
	lines := strings.Split(src, "\n")
	endOfFunction := FindEndOfFunction(lines, 1)
	assert.Equal(t, 7, endOfFunction)
}

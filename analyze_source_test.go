package erro

import (
	"github.com/stretchr/testify/assert"
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
	lastWrite := lastWriteToVar(src, "y")
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

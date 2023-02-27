package erro

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestFindStartEnd(t *testing.T) {
	failingLine := "err := calc(x,y)"
	i0 := strings.Index(failingLine, "(")
	i1 := i0
	start, end := findStartEnd(failingLine, i0, i1)

	assert.Equal(t, 11, start)
	assert.Equal(t, 15, end)

}

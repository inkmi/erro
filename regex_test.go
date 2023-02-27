package erro

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestFindStartEnding(t *testing.T) {
	failingLine := "err := calc(x,y)"
	end := findStartEnd2(failingLine, strings.Index(failingLine, "("))
	assert.Equal(t, 15, end)
}

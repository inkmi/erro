package internal

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindStartEnding(t *testing.T) {
	failingLine := "err := calc(x,y)"
	end := findStartEnd(failingLine, strings.Index(failingLine, "("))
	assert.Equal(t, 15, end)
}

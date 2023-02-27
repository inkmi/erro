package internal

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestFindStartEnding(t *testing.T) {
	failingLine := "err := calc(x,y)"
	end := findStartEnd(failingLine, strings.Index(failingLine, "("))
	assert.Equal(t, 15, end)
}

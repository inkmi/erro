package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShortFilePath(t *testing.T) {
	t.Parallel()
	exp := "error.go"
	assert.Equal(t, exp, getShorterFilePath("/stephan/src/error.go", "/stephan/src/"))
}

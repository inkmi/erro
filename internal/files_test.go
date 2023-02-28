package internal

import (
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShortFilePath(t *testing.T) {
	t.Parallel()
	exp := "error.go"
	assert.Equal(t, exp, getShorterFilePath("/stephan/src/error.go", "/stephan/src/"))
}

func TestReadSource(t *testing.T) {
	appFS := afero.NewMemMapFs()
	// create test files and directories
	afero.WriteFile(appFS, "error.go", []byte("x = 4"), 0644)
	source := ReadSourceFs("error.go", appFS)
	assert.Equal(t, 1, len(source))
	assert.Equal(t, "x = 4", source[0])
}

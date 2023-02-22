package erro

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorfIsSource(t *testing.T) {
	se := errors.New("file not found")
	ne := Errorf("Error: %w", se)
	assert.True(t, errors.Is(ne, se))
}

func TestNewIsSource(t *testing.T) {
	se := errors.New("file not found")
	ne := New("Can't do", se)
	assert.True(t, errors.Is(ne, se))
}

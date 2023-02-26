package erro

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMin(t *testing.T) {
	exp := 3
	assert.Equal(t, exp, min(3, 5))
}

func TestMinMinus(t *testing.T) {
	exp := -1
	assert.Equal(t, exp, min(3, -1))
}

func TestMax(t *testing.T) {
	exp := 5
	assert.Equal(t, exp, max(3, 5))
}

func TestMaxMinus(t *testing.T) {
	exp := 3
	assert.Equal(t, exp, max(3, -1))
}

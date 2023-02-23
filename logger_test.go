package erro

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDiff(t *testing.T) {
	s1 := []string{"a", "b", "c"}
	s2 := []string{"b", "d", "e"}
	res := diff(s1, s2)
	exp := []string{"a", "c"}
	assert.Equal(t, exp, res)
}

func TestContains(t *testing.T) {
	s := []string{"a", "b", "c"}
	assert.True(t, contains(s, "a"))
	assert.False(t, contains(s, "z"))
}

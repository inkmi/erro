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

func TestDiffEmpty(t *testing.T) {
	s1 := []string{"a", "b", "c"}
	var s2 []string
	res := diff(s1, s2)
	exp := []string{"a", "b", "c"}
	assert.Equal(t, exp, res)
}

func TestDiffFirstEmpty(t *testing.T) {
	var s1 []string
	s2 := []string{"a"}
	res := diff(s1, s2)
	var exp []string
	assert.Equal(t, exp, res)
}

func TestContains(t *testing.T) {
	s := []string{"a", "b", "c"}
	assert.True(t, contains(s, "a"))
	assert.False(t, contains(s, "z"))
}

func TestContainsEmpty(t *testing.T) {
	var s []string
	assert.False(t, contains(s, "z"))
}

func TestContainsInt(t *testing.T) {
	s := []int{1, 2, 3}
	assert.True(t, isIntInSlice(1, s))
	assert.False(t, isIntInSlice(4, s))
}

func TestContainsIntEmpty(t *testing.T) {
	s := []int{}
	assert.False(t, isIntInSlice(4, s))
}

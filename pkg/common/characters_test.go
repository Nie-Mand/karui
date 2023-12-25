package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAlpha(t *testing.T) {
	assert.True(t, IsAlpha("a"))
	assert.True(t, IsAlpha("A"))
	assert.True(t, IsAlpha("z"))
	assert.True(t, IsAlpha("Z"))
	assert.True(t, IsAlpha("abc"))
	assert.True(t, IsAlpha("ABC"))
	assert.True(t, IsAlpha("abcABC"))
	assert.False(t, IsAlpha("abcABC123"))
	assert.False(t, IsAlpha("abcABC123!@#"))
	assert.False(t, IsAlpha(""))
}

func TestIsNumeric(t *testing.T) {
	assert.True(t, IsNumeric("0"))
	assert.True(t, IsNumeric("1"))
	assert.True(t, IsNumeric("9"))
	assert.True(t, IsNumeric("123"))
	assert.False(t, IsNumeric("123abc"))
	assert.False(t, IsNumeric("abc"))
	assert.False(t, IsNumeric(""))
}

func TestIsAlphanumeric(t *testing.T) {
	assert.True(t, IsAlphanumeric("a"))
	assert.True(t, IsAlphanumeric("A"))
	assert.True(t, IsAlphanumeric("z"))
	assert.True(t, IsAlphanumeric("Z"))
	assert.True(t, IsAlphanumeric("abc"))
	assert.True(t, IsAlphanumeric("ABC"))
	assert.True(t, IsAlphanumeric("abcABC"))
	assert.True(t, IsAlphanumeric("abcABC123"))
	assert.False(t, IsAlphanumeric("abcABC123!@#"))
	assert.False(t, IsAlphanumeric(""))
	assert.True(t, IsAlphanumeric("0"))
	assert.True(t, IsAlphanumeric("1"))
	assert.True(t, IsAlphanumeric("9"))
	assert.True(t, IsAlphanumeric("123"))
	assert.True(t, IsAlphanumeric("123abc"))
}
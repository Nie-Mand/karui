package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStringIterator(t *testing.T) {
	source := ""
	it := NewStringIterator(source)
	assert.NotNil(t, it)
}

func TestPeek(t *testing.T) {
	source := "a"
	it := NewStringIterator(source)
	peek, err := it.Peek()
	assert.Equal(t, "a", peek)
	assert.Nil(t, err)
}

func TestNext(t *testing.T) {
	source := "a"
	it := NewStringIterator(source)
	next, err := it.Next()
	assert.Equal(t, "a", next)
	assert.Nil(t, err)
}

func TestHasNext(t *testing.T) {
	source := "a"
	it := NewStringIterator(source)
	assert.True(t, it.HasNext())
}

func TestHasRemaining(t *testing.T) {
	source := "a"
	it := NewStringIterator(source)
	assert.True(t, it.HasRemaining(1))
	assert.False(t, it.HasRemaining(2))
}

func TestReset(t *testing.T) {
	source := "abc"
	it := NewStringIterator(source)
	it.Next()
	p, _ := it.Peek()
	assert.Equal(t, "b", p)
	it.Reset()
	p, _ = it.Peek()
	assert.Equal(t, "a", p)
}
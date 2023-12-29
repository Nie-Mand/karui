package tokens

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestNewTokenIterator(t *testing.T) {
	tokens := []Token{}
	it := NewTokenIterator(tokens)
	assert.NotNil(t, it)
}

func TestPeek(t *testing.T) {
	tokens := []Token{
		{ Type: Identifier, Value: "a" },
	}
	it := NewTokenIterator(tokens)
	peek, err := it.Peek()
	assert.Equal(t, "a", peek.Value)
	assert.Nil(t, err)
}

func TestNext(t *testing.T) {
	tokens := []Token{
		{ Type: Identifier, Value: "a" },
	}
	it := NewTokenIterator(tokens)
	next, err := it.Next()
	assert.Equal(t, "a", next.Value)
	assert.Nil(t, err)
}

func TestHasNext(t *testing.T) {
	tokens := []Token{
		{ Type: Identifier, Value: "a" },
	}
	it := NewTokenIterator(tokens)
	assert.True(t, it.HasNext())
}

func TestReset(t *testing.T) {
	tokens := []Token{
		{ Type: Identifier, Value: "a" },
		{ Type: Identifier, Value: "b" },
		{ Type: Identifier, Value: "c" },
	}
	it := NewTokenIterator(tokens)
	it.Next()
	p, _ := it.Peek()
	assert.Equal(t, "b", p.Value)
	it.Reset()
	p, _ = it.Peek()
	assert.Equal(t, "a", p.Value)
}
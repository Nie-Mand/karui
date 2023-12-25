package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	source := ""
	lexer  := NewLexer(source)
	assert.NotNil(t, lexer)
}

func TestPeek(t *testing.T) {
	source := "a"
	lexer  := NewLexer(source)
	assert.Equal(t, "a", lexer.Peek())
}

func TestPeekOffset(t *testing.T) {
	source := "abc"
	lexer  := NewLexer(source)
	assert.Equal(t, "c", lexer.PeekOffset(2))
}

func TestConsume(t *testing.T) {
	source := "ab"
	lexer  := NewLexer(source)
	c, _ := lexer.Consume()
	assert.Equal(t, "a", c)
	assert.Equal(t, "b", lexer.Peek())
}

func TestLexerize(t *testing.T) {
	source := "a"
	lexer  := NewLexer(source)
	tokens, _ := lexer.Lexerize()
	assert.Equal(t, []Token{{ Type: Ident, Value: "a" }}, tokens)

	source = "let a = 1;"
	lexer  = NewLexer(source)
	tokens, _ = lexer.Lexerize()
	assert.Equal(t, []Token{
		{ Type: Let },
		{ Type: Ident, Value: "a" },
		{ Type: Equal },
		{ Type: IntIdent, Value: "1" },
		{ Type: Semicolon },
	}, tokens)
}
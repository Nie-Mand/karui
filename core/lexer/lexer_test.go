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
package lexer

import (
	"Nie-Mand/karui/core/internal/tokens"
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
	_tokens, _ := lexer.Lexerize()
	assert.Equal(t, []tokens.Token{{ Type: tokens.Identifier, Value: "a" }}, _tokens)

	source = "let a = 1;"
	lexer  = NewLexer(source)
	_tokens, _ = lexer.Lexerize()
	assert.Equal(t, []tokens.Token{
		{ Type: tokens.Let },
		{ Type: tokens.Identifier, Value: "a" },
		{ Type: tokens.Equal },
		{ Type: tokens.IntLiteral, Value: "1" },
		{ Type: tokens.Semicolon },
	}, _tokens)
}
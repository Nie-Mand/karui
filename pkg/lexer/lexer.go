package lexer

import (
	"Nie-Mand/karui/pkg/common"
	"strings"
)

type Lexer struct {
	index  int
	source string
}

func NewLexer(source string) *Lexer {
	return &Lexer{
		index: 0, 
		source: strings.Join(strings.Split(source, "\n"), SEPERATOR),
	}
}

func (l *Lexer) Lexerize() ([]Token, error) {
	tokens := []Token{}
	buffer := ""

	for l.Peek() != "" {
		character := l.Peek()

		if common.IsAlpha(character) {
			c, _ := l.Consume()
			buffer += c
			for common.IsAlphanumeric(l.Peek()) {
				c, more := l.Consume()
				buffer += c
				if !more {
					break
				} 
			}

			switch buffer {
			case EXIT:
				tokens = append(tokens, Token{ Type: Exit })
			case LET:
				tokens = append(tokens, Token{ Type: Let })
			case PUTS:
				tokens = append(tokens, Token{ Type: Puts })
			default:
				tokens = append(tokens, Token{ Type: Ident, Value: buffer })
			}

			buffer = ""
		} else if common.IsNumeric(character) {
			c, _ := l.Consume()
			buffer += c
			for common.IsNumeric(l.Peek()) {
				c, more := l.Consume()
				buffer += c

				if !more {
					break
				} 
			}
			tokens = append(tokens, Token{IntIdent, buffer})
			buffer = ""
		} else {
			switch character {
			case SEMICOLON:
				tokens = append(tokens, Token{ Type: Semicolon })
			case LEFT_PARENTHESIS:
				tokens = append(tokens, Token{ Type: LeftParenthesis })
			case RIGHT_PARENTHESIS:
				tokens = append(tokens, Token{ Type: RightParenthesis })
			case EQUAL:
				tokens = append(tokens, Token{ Type: Equal })
			case PLUS:
				tokens = append(tokens, Token{ Type: Plus })
			case MINUS:
				tokens = append(tokens, Token{ Type: Minus })
			case MULTIPLY:
				tokens = append(tokens, Token{ Type: Multiply })
			case DIVIDE:
				tokens = append(tokens, Token{ Type: Divide })
			case MODULO:
				tokens = append(tokens, Token{ Type: Modulo })
			case " ":
				// ignore
			case "\t":
				// ignore
			default:
				return nil, ErrInvalidCharacter
			}
			l.Consume()
		}
	}
	
	l.index = 0
	return tokens, nil
}

func (l *Lexer) Peek() string {
	return l.PeekOffset(0)
}

func (l *Lexer) PeekOffset(offset int) string {
	if l.index + offset < len(l.source) {
		return string(l.source[l.index + offset])
	}
	return ""
}

func (l *Lexer) Consume() (string, bool) {
	if l.index < len(l.source) {
		char := string(l.source[l.index])
		l.index++
		return char, true
	}

	return "", false
}

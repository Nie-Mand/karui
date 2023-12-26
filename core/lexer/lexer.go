package lexer

import (
	"Nie-Mand/karui/core/common"
	"Nie-Mand/karui/core/utils"
	"strings"
)

type Lexer struct {
	iterator common.Iterator[string]
}

func NewLexer(source string) *Lexer {
	parsed := strings.Join(strings.Split(source, "\n"), SEPERATOR)
	return &Lexer{
		iterator: utils.NewStringIterator(parsed),
	}
}

func (l *Lexer) Lexerize() ([]Token, error) {
	tokens := []Token{}
	buffer := ""

	for l.iterator.HasNext() {
		if p, err := l.iterator.Peek(); err == nil && utils.IsAlpha(p) {
			for {
				if p, err := l.iterator.Peek(); err == nil && utils.IsAlphanumeric(p) {
					c, _ := l.iterator.Next()
					buffer += c
				} else {
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
			case IF:
				tokens = append(tokens, Token{ Type: If })
			default:
				tokens = append(tokens, Token{ Type: Identifier, Value: buffer })
			}

			buffer = ""
		} else if err == nil && utils.IsNumeric(p) {
			for {
				if p, err := l.iterator.Peek(); err == nil && utils.IsNumeric(p) {
					c, _ := l.iterator.Next()
					buffer += c
				} else {
					break
				}
			}

			tokens = append(tokens, Token{IntLiteral, buffer})
			buffer = ""
		} else {
			switch p {
			case SEMICOLON:
				tokens = append(tokens, Token{ Type: Semicolon })
			case OPEN_PARENTHESIS:
				tokens = append(tokens, Token{ Type: OpenParenthesis })
			case CLOSE_PARENTHESIS:
				tokens = append(tokens, Token{ Type: CloseParenthesis })
			case OPEN_CURLY:
				tokens = append(tokens, Token{ Type: OpenCurly })
			case CLOSE_CURLY:
				tokens = append(tokens, Token{ Type: CloseCurly })
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
			l.iterator.Next()
		}
	}
	
	l.iterator.Reset()
	return tokens, nil
}

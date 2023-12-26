package lexer

import (
	"Nie-Mand/karui/core/common"
	"Nie-Mand/karui/core/lexer/tokens"
	"Nie-Mand/karui/core/utils"
	"strings"
)

type Lexer struct {
	iterator common.Iterator[string]
}

func NewLexer(source string) *Lexer {
	parsed := strings.Join(strings.Split(source, "\n"), SEPERATOR.String())
	return &Lexer{
		iterator: utils.NewStringIterator(parsed),
	}
}

func (l *Lexer) Lexerize() ([]tokens.Token, error) {
	_tokens := []tokens.Token{}
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
				_tokens = append(_tokens, tokens.Token{ Type: tokens.Exit })
			case LET:
				_tokens = append(_tokens, tokens.Token{ Type: tokens.Let })
			case PUTS:
				_tokens = append(_tokens, tokens.Token{ Type: tokens.Puts })
			case IF:
				_tokens = append(_tokens, tokens.Token{ Type: tokens.If })
			default:
				_tokens = append(_tokens, tokens.Token{ Type: tokens.Identifier, Value: buffer })
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

			_tokens = append(_tokens, tokens.Token{
				Type: tokens.IntLiteral,
				Value: buffer,
			})
			buffer = ""
		} else {
			switch p {
			case SEMICOLON:
				_tokens = append(_tokens, tokens.Token{ Type: tokens.Semicolon })
			case OPEN_PARENTHESIS:
				_tokens = append(_tokens, tokens.Token{ Type: tokens.OpenParenthesis })
			case CLOSE_PARENTHESIS:
				_tokens = append(_tokens, tokens.Token{ Type: tokens.CloseParenthesis })
			case OPEN_CURLY:
				_tokens = append(_tokens, tokens.Token{ Type: tokens.OpenCurly })
			case CLOSE_CURLY:
				_tokens = append(_tokens, tokens.Token{ Type: tokens.CloseCurly })
			case EQUAL:
				_tokens = append(_tokens, tokens.Token{ Type: tokens.Equal })
			case PLUS:
				_tokens = append(_tokens, tokens.Token{ Type: tokens.Plus })
			case MINUS:
				_tokens = append(_tokens, tokens.Token{ Type: tokens.Minus })
			case MULTIPLY:
				_tokens = append(_tokens, tokens.Token{ Type: tokens.Multiply })
			case DIVIDE:
				_tokens = append(_tokens, tokens.Token{ Type: tokens.Divide })
			case MODULO:
				_tokens = append(_tokens, tokens.Token{ Type: tokens.Modulo })
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
	return _tokens, nil
}

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
	lines := strings.Split(source, "\n")
	parsed := []string{}

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if line[0] == COMMENT.String()[0] {
			continue
		}

		if line[len(line) - 1] != ';' && line[len(line) - 1] != '{' && line[len(line) - 1] != '}' {
			line += ";"
		}

		parsed = append(parsed, line)
	}

	return &Lexer{
		iterator: utils.NewStringIterator(strings.Join(parsed, "")),
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
			case tokens.Exit.String():
				_tokens = append(_tokens, tokens.Token{ Type: tokens.Exit })
			case tokens.Let.String():
				_tokens = append(_tokens, tokens.Token{ Type: tokens.Let })
			case tokens.Puts.String():
				_tokens = append(_tokens, tokens.Token{ Type: tokens.Puts })
			case tokens.If.String():
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
			case tokens.Semicolon.String():
				_tokens = append(_tokens, tokens.Token{ Type: tokens.Semicolon })
			case tokens.OpenParenthesis.String():
				_tokens = append(_tokens, tokens.Token{ Type: tokens.OpenParenthesis })
			case tokens.CloseParenthesis.String():
				_tokens = append(_tokens, tokens.Token{ Type: tokens.CloseParenthesis })
			case tokens.OpenCurly.String():
				_tokens = append(_tokens, tokens.Token{ Type: tokens.OpenCurly })
			case tokens.CloseCurly.String():
				_tokens = append(_tokens, tokens.Token{ Type: tokens.CloseCurly })
			case tokens.Equal.String():
				_tokens = append(_tokens, tokens.Token{ Type: tokens.Equal })
			case tokens.Plus.String():
				_tokens = append(_tokens, tokens.Token{ Type: tokens.Plus })
			case tokens.Minus.String():
				_tokens = append(_tokens, tokens.Token{ Type: tokens.Minus })
			case tokens.Multiply.String():
				_tokens = append(_tokens, tokens.Token{ Type: tokens.Multiply })
			case tokens.Divide.String():
				_tokens = append(_tokens, tokens.Token{ Type: tokens.Divide })
			case tokens.Modulo.String():
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

package lexer

import (
	"github.com/Nie-Mand/karui/pkg/utils"
)


type Lexer[Token any] struct {
	iterator utils.Iterator[string]
	tokinzers []TokenParser[Token]
}

func NewLexer[Token any](formatter Formatter, source string) *Lexer[Token] {
	source = formatter.Format(source)
	return &Lexer[Token]{
		iterator: utils.NewStringIterator(source),
		tokinzers: []TokenParser[Token]{},
	}
}

func (l *Lexer[Token]) AddTokenParser(parser TokenParser[Token]) {
	l.tokinzers = append(l.tokinzers, parser)
}

func (l *Lexer[Token]) Lexerize() ([]Token, error) {
	_tokens := []Token{}

	for l.iterator.HasNext() {
		p, _ := l.iterator.Peek()

		for _, tokenizer := range l.tokinzers {
			if tokenizer.StartsWidth(p) {
				token, err := tokenizer.Parse(l.iterator)
				if err != nil{
					return nil, err
				}

				if token == nil {
					continue
				}

				_tokens = append(_tokens, *token)
				break
			}
		}
	}
	
	l.iterator.Reset()
	return _tokens, nil
}

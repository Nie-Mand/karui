package tokenizers

import (
	"github.com/Nie-Mand/karui/core/internal/tokens"
	"github.com/Nie-Mand/karui/pkg/utils"
)

type OperatorsTokenizer struct {}

func (st *OperatorsTokenizer) Parse(it utils.Iterator[string]) (*tokens.Token, error) {
	p, _ := it.Peek()

	token := &tokens.Token{}
	switch p {
	case tokens.Equal.String():
		token.Type = tokens.Equal
	case tokens.Plus.String():
		token.Type = tokens.Plus
	case tokens.Minus.String():
		token.Type = tokens.Minus
	case tokens.Multiply.String():
		token.Type = tokens.Multiply
	case tokens.Divide.String():
		token.Type = tokens.Divide
	case tokens.Modulo.String():
		token.Type = tokens.Modulo
	default:
		return nil, nil
	}

	it.Next()
	return token, nil
}

func (st *OperatorsTokenizer) StartsWidth(p string) bool {
	return p == tokens.Equal.String() || p == tokens.Plus.String() || p == tokens.Minus.String() || p == tokens.Multiply.String() || p == tokens.Divide.String() || p == tokens.Modulo.String()
}
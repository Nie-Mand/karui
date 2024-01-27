package tokenizers

import (
	"github.com/Nie-Mand/karui/core/internal/tokens"
	"github.com/Nie-Mand/karui/pkg/utils"
)

type BracketsTokenizer struct {}

func (st *BracketsTokenizer) Parse(it utils.Iterator[string]) (*tokens.Token, error) {
	p, _ := it.Peek()

	token := &tokens.Token{}
	switch p {
	case tokens.OpenParenthesis.String():
		token.Type = tokens.OpenParenthesis
	case tokens.CloseParenthesis.String():
		token.Type = tokens.CloseParenthesis
	case tokens.OpenCurly.String():
		token.Type = tokens.OpenCurly
	case tokens.CloseCurly.String():
		token.Type = tokens.CloseCurly
	default:
		return nil, nil
	}

	it.Next()
	return token, nil
}

func (st *BracketsTokenizer) StartsWidth(p string) bool {
	return p == tokens.OpenParenthesis.String() || p == tokens.CloseParenthesis.String() || p == tokens.OpenCurly.String() || p == tokens.CloseCurly.String()
}
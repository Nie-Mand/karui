package tokenizers

import (
	"Nie-Mand/karui/core/internal/tokens"
	"Nie-Mand/karui/pkg/utils"
)

type NumbersTokenizer struct{}

func (st *NumbersTokenizer) Parse(it utils.Iterator[string]) (*tokens.Token, error) {
	buffer := ""

	for {
		if p, err := it.Peek(); err == nil && utils.IsNumeric(p) {
			c, _ := it.Next()
			buffer += c
		} else {
			break
		}
	}

	return &tokens.Token{Type: tokens.IntLiteral, Value: buffer}, nil
}

func (st *NumbersTokenizer) StartsWidth(p string) bool {
	return utils.IsNumeric(p)
}

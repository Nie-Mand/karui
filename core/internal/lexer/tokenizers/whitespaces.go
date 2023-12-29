package tokenizers

import (
	"Nie-Mand/karui/core/internal/tokens"
	"Nie-Mand/karui/pkg/utils"
)

type WhitespacesTokenizer struct {}

func (st *WhitespacesTokenizer) Parse(it utils.Iterator[string]) (*tokens.Token, error) {
	it.Next()
	return nil, nil
}

func (st *WhitespacesTokenizer) StartsWidth(p string) bool {
	return p == " " || p == "\t"
}
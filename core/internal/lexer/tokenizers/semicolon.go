package tokenizers

import (
	"Nie-Mand/karui/core/internal/tokens"
	"Nie-Mand/karui/pkg/utils"
)

type SemicolonTokenizer struct {}

func (st *SemicolonTokenizer) Parse(it utils.Iterator[string]) (*tokens.Token, error) {
	it.Next()
	return &tokens.Token{ Type: tokens.Semicolon }, nil
}

func (st *SemicolonTokenizer) StartsWidth(p string) bool {
	return p == tokens.Semicolon.String()
}
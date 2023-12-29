package tokenizers

import (
	"Nie-Mand/karui/core/internal/tokens"
	"Nie-Mand/karui/pkg/utils"
)

type WordsTokenizer struct {}

func (st *WordsTokenizer) Parse(it utils.Iterator[string]) (*tokens.Token, error) {
	buffer := ""

	for {
		if p, err := it.Peek(); err == nil && utils.IsAlphanumeric(p) {
			c, _ := it.Next()
			buffer += c
		} else {
			break
		}
	}

	switch buffer {
	case tokens.Exit.String():
		return &tokens.Token{ Type: tokens.Exit }, nil
	case tokens.Let.String():
		return &tokens.Token{ Type: tokens.Let }, nil
	case tokens.Puts.String():
		return &tokens.Token{ Type: tokens.Puts }, nil
	case tokens.If.String():
		return &tokens.Token{ Type: tokens.If }, nil
	default:
		return &tokens.Token{ Type: tokens.Identifier, Value: buffer }, nil
	}
}

func (st *WordsTokenizer) StartsWidth(p string) bool {
	return utils.IsAlpha(p)
}
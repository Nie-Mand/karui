package lexer

import "github.com/Nie-Mand/karui/pkg/utils"

type TokenParser[Token any] interface {

	StartsWidth(string) bool
	
	Parse(utils.Iterator[string]) (*Token, error)
}
package lexer

import "errors"

var (
	ErrInvalidCharacter = errors.New("invalid character")
	ErrSourceEmpty = errors.New("source is empty")
)
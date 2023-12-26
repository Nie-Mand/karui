package lexer

import "Nie-Mand/karui/core/common"

type TokenIterator struct {
	index  int
	base []Token
}

func NewTokenIterator(tokens []Token) common.Iterator[Token] {
	return &TokenIterator{
		index: 0, 
		base: tokens,
	}
}

func (it *TokenIterator) Peek() (Token, error) {
	if it.HasNext() {
		return it.base[it.index], nil
	}

	return Token{}, common.ErrIteratorOutOfRange
}

func (it *TokenIterator) Next() (Token, error) {
	if it.HasNext() {
		it.index++
		return it.base[it.index - 1], nil
	}
	
	return Token{}, common.ErrIteratorOutOfRange
}

func (it *TokenIterator) HasNext() bool {
	return it.HasRemaining(1)
}

func (it *TokenIterator) HasRemaining(offset int) bool {
	return it.index + offset <= len(it.base)
}

func (it *TokenIterator) Reset() {
	it.index = 0
}
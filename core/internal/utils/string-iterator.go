package utils

import "Nie-Mand/karui/core/common"

type StringIterator struct {
	index  int
	base string
}

func NewStringIterator(source string) common.Iterator[string] {
	return &StringIterator{
		index: 0, 
		base: source,
	}
}

func (it *StringIterator) Peek() (string, error) {
	if it.HasNext() {
		return string(it.base[it.index]), nil
	}

	return "", common.ErrIteratorOutOfRange
}

func (it *StringIterator) Next() (string, error) {
	if it.HasNext() {
		it.index++
		return string(it.base[it.index - 1]), nil
	}
	
	return "", common.ErrIteratorOutOfRange
}

func (it *StringIterator) HasNext() bool {
	return it.HasRemaining(1)
}

func (it *StringIterator) HasRemaining(offset int) bool {
	return it.index + offset <= len(it.base)
}

func (it *StringIterator) Reset() {
	it.index = 0
}
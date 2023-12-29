package utils

import "errors"


var (
	ErrIteratorOutOfRange = errors.New("Iterator out of range")
)

type Iterator[IT any] interface {

	Peek() (IT, error)

	Next() (IT, error)

	HasNext() bool
	
	Reset()
}
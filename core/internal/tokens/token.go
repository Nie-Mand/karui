package tokens

type Token struct {
	Type  TokenType
	Value string
}

func (t *Token) BinaryPrec() int {
	return -1
} 
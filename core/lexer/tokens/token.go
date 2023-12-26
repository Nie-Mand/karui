package tokens

type Token struct {
	Type  TokenType
	Value string
}

func (t *Token) BinaryPrec() int {
	if t.Type == Plus || t.Type == Minus {
		return 0
	}

	if t.Type == Multiply || t.Type == Divide {
		return 1
	}

	return -1
} 
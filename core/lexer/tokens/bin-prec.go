package tokens

func BinaryPrec(token Token) int {
	// - binary precision: int
	//   - if its -+, then its 0
	//   - if its */, then its 1
	//   - else its null

	// if token.Type == Plus || token.Type == Minus {
	// 	return 0
	// } else if token.Type == Multiply || token.Type == Divide {
	// 	return 1
	// } else {
	// 	return -1
	// }
	return -1
}
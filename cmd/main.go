package main

import (
	"Nie-Mand/karui/core/lexer"
	"fmt"
)

func main() {
	source := "let a = 1;"
	lexer  := lexer.NewLexer(source)
	tokens, err := lexer.Lexerize()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(tokens)
	}
}
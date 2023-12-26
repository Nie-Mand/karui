package main

import (
	"Nie-Mand/karui/core/lexer"
	"Nie-Mand/karui/core/parser"
	"fmt"
)

func main() {
	source := "let a = 1;"
	lexer  := lexer.NewLexer(source)
	tokens, err := lexer.Lexerize()
	if err != nil {
		panic(err)
	} else {
		fmt.Println(tokens)
	}

	parser := parser.NewParser(tokens)
	program, err := parser.ParseProgram()
	if err != nil {
		panic(err)
	} else {
		fmt.Println(program)
	}
}
package main

import (
	"Nie-Mand/karui/core/lexer"
	"Nie-Mand/karui/core/parser"
	"fmt"
	"os"
)

func main() {
	source, err := os.ReadFile("examples/1.kui")
	if err != nil {
		panic(err)
	}
	
	lexer  := lexer.NewLexer(string(source))
	tokens, err := lexer.Lexerize()
	if err != nil {
		panic(err)
	}

	parser := parser.NewParser(tokens)
	program, err := parser.ParseProgram()
	if err != nil {
		panic(err)
	} else {
		fmt.Println(program)
	}
}
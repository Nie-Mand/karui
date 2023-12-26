package main

import (
	"Nie-Mand/karui/core/executor"
	"Nie-Mand/karui/core/lexer"
	"Nie-Mand/karui/core/parser"
	"fmt"
	"os"
)

func main() {
	fmt.Println("[*] Executing examples/2.kui")
	source, err := os.ReadFile("examples/2.kui")
	if err != nil {
		panic(err)
	}

	lexer := lexer.NewLexer(string(source))
	tokens, err := lexer.Lexerize()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("[*] Successfully lexerized program")
	}

	parser := parser.NewParser(tokens)
	program, err := parser.ParseProgram()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("[*] Successfully parsed program")
	}

	executor := executor.NewExecutor(program)

	err = executor.ExecuteProgram()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("[*] Successfully executed program")
	}

}

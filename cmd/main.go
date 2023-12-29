package main

import (
	// "Nie-Mand/karui/core/internal/lexer"
	"Nie-Mand/karui/core/execute"
	"os"
)

func main() {
	source, err := os.ReadFile("examples/2.kui")
	if err != nil {
		panic(err)
	}

	execute.Execute(string(source))
	

	// parser := parser.NewParser(tokens)
	// program, err := parser.ParseProgram()
	// if err != nil {
	// 	panic(err)
	// }

	// executor := executor.NewExecutor(program)

	// err = executor.ExecuteProgram()
	// if err != nil {
	// 	panic(err)
	// }
}

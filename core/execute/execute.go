package execute

import (
	// "Nie-Mand/karui/core/lexer"
	// "Nie-Mand/karui/core/internal/lexer"

	"Nie-Mand/karui/core/internal/lexer"
	"fmt"
)


func Execute(source string) {
	// lexer := lexer.NewLexer(string(source))
	lexer := lexer.NewKaruiLexer(string(source))
	tokens, err := lexer.Lexerize()
	if err != nil {
		panic(err)
	} 

	fmt.Println(tokens)	
}
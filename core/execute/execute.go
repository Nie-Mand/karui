package execute

import (
	"github.com/Nie-Mand/karui/core/internal/lexer"

	"fmt"
)

func Execute(source string) {
	karuiLexer := lexer.NewKaruiLexer(string(source))
	tokens, err := karuiLexer.Lexerize()
	if err != nil {
		panic(err)
	}

	fmt.Println(tokens)
}

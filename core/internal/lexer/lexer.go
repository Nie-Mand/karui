package lexer

import (
	"github.com/Nie-Mand/karui/core/internal/lexer/tokenizers"
	"github.com/Nie-Mand/karui/core/internal/tokens"
	"github.com/Nie-Mand/karui/core/internal/utils"
	"github.com/Nie-Mand/karui/pkg/lexer"
)

type KaruiLexer struct {
	Lexer lexer.Lexer[tokens.Token]
}

func NewKaruiLexer(source string) *KaruiLexer {
	formatter := utils.BasicFormatter{}
	lexer := lexer.NewLexer[tokens.Token](&formatter, source)

	// Register Tokenizers
	lexer.AddTokenParser(&tokenizers.WordsTokenizer{})
	lexer.AddTokenParser(&tokenizers.NumbersTokenizer{})
	lexer.AddTokenParser(&tokenizers.SemicolonTokenizer{})
	lexer.AddTokenParser(&tokenizers.BracketsTokenizer{})
	lexer.AddTokenParser(&tokenizers.OperatorsTokenizer{})
	lexer.AddTokenParser(&tokenizers.WhitespacesTokenizer{})

	return &KaruiLexer{
		Lexer: *lexer,
	}
}

func (kl *KaruiLexer) Lexerize() ([]tokens.Token, error) {
	return kl.Lexer.Lexerize()
}
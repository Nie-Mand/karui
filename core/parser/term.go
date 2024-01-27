package parser

import (
	"fmt"

	"github.com/Nie-Mand/karui/core/internal/tokens"
)

type TermType interface {
	IsTerm()
	String() string
	Evaluate(*map[string]int) (int, error)
}

// IdentifierTerm
type IdentifierTerm struct {
	Token tokens.Token
}

func (IdentifierTerm) IsTerm() {}
func (t *IdentifierTerm) Evaluate(memory *map[string]int) (int, error) {

	if memory == nil {
		return 0, ErrMemoryNotInitialized
	}

	if value, exists := (*memory)[t.Token.Value]; exists {
		return value, nil
	}
	
	return 0, ErrVariableNotInitialized
}

func (t *IdentifierTerm) String() string {
	return t.Token.Value
}

// IntLiteralTerm
type IntLiteralTerm struct {
	Token tokens.Token
}

func (IntLiteralTerm) IsTerm() {}
func (t *IntLiteralTerm) Evaluate(_ *map[string]int) (int, error) {
	var value int
	_, err := fmt.Sscanf(t.Token.Value, "%d", &value)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (t *IntLiteralTerm) String() string {
	return t.Token.Value
}

// ParenthesizedTerm
type ParenthesizedTerm struct {
	Expression ExpressionType
}

func (ParenthesizedTerm) IsTerm() {}
func (t *ParenthesizedTerm) Evaluate(memory *map[string]int) (int, error) {
	return t.Expression.Evaluate(memory)
}

func (t *ParenthesizedTerm) String() string {
	return "(" + t.Expression.String() + ")"
}
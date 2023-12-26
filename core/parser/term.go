package parser

import "Nie-Mand/karui/core/lexer/tokens"

type TermType interface {
	IsTerm()
	IsExpression()
	String() string
}

// IdentifierTerm
type IdentifierTerm struct {
	Token tokens.Token
}

func (IdentifierTerm) IsTerm() {}
func (IdentifierTerm) IsExpression() {}
func (i *IdentifierTerm) String() string {
	return i.Token.Value
}


// IntLiteralTerm
type IntLiteralTerm struct {
	Token tokens.Token
}

func (IntLiteralTerm) IsTerm() {}
func (IntLiteralTerm) IsExpression() {}
func (i *IntLiteralTerm) String() string {
	return i.Token.Value
}


// ParenthesizedTerm
type ParenthesizedTerm struct {
	Expression ExpressionType
}

func (ParenthesizedTerm) IsTerm() {}
func (ParenthesizedTerm) IsExpression() {}
func (p *ParenthesizedTerm) String() string {
	return "(" + p.Expression.String() + ")"
}
package parser

import (
	"Nie-Mand/karui/core/lexer/tokens"
)


type StatementType interface {
	IsStatement()
	String() string
}

// PutStatement
type PutStatement struct {
	Expression ExpressionType
}

func (PutStatement) IsStatement() {}

func (PutStatement) String() string {
	return "PutStatement"
}


// LetStatement
type LetStatement struct {
	Token tokens.Token
	Expression ExpressionType
}

func (LetStatement) IsStatement() {}

func (l *LetStatement) String() string {
	representation := l.Token.Value + " <- " + l.Expression.String()
	return representation
}


// IfStatement
type IfStatement struct {
	Condition ExpressionType
	Scope Scope
}

func (IfStatement) IsStatement() {}

func (IfStatement) String() string {
	return "IfStatement"
}


// ExitStatement
type ExitStatement struct {
	ExitCode ExpressionType
}

func (ExitStatement) IsStatement() {}

func (ExitStatement) String() string {
	return "ExitStatement"
}
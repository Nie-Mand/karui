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

func (p * PutStatement) String() string {
	representation := "puts " + p.Expression.String()
	return representation
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

func (s *IfStatement) String() string {
	representation := "if: " + s.Condition.String() + " do: \n"
	for _, statement := range s.Scope.Statements {
		representation += "\t" + statement.String() + "\n"
	}

	return representation
}


// ExitStatement
type ExitStatement struct {
	ExitCode ExpressionType
}

func (ExitStatement) IsStatement() {}

func (s *ExitStatement) String() string {
	return "exit with error code equals to " + s.ExitCode.String()
}
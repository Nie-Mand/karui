package parser

import (
	"Nie-Mand/karui/core/lexer/tokens"
	"fmt"
)


type StatementType interface {
	Execute() error
}

// PutStatement
type PutStatement struct {
	Expression ExpressionType
}

func (p *PutStatement) Execute() error {
	fmt.Println(p.Expression.String())
	return nil
}

// LetStatement
type LetStatement struct {
	Token tokens.Token
	Expression ExpressionType
}

func (LetStatement) Execute() error {
	// TODO: Implement
	return nil
}

// IfStatement
type IfStatement struct {
	Condition ExpressionType
	Scope Scope
}

func (s *IfStatement) Execute() error {
	// evaluation, err := s.Condition.Evaluate()
	// TODO: Implement
	fmt.Println("Evaluating condition: " + s.Condition.String())

	for _, statement := range s.Scope.Statements {
		err := statement.Execute()
		if err != nil {
			return err
		}
	}

	return nil
}


// ExitStatement
type ExitStatement struct {
	ExitCode ExpressionType
}

func (s *ExitStatement) Execute() error {
	return ErrProgramExits(s.ExitCode.String())
}
package parser

import (
	"fmt"

	"github.com/Nie-Mand/karui/core/internal/tokens"
)


type StatementType interface {
	Execute(*map[string]int) error
}

// PutStatement
type PutStatement struct {
	Expression ExpressionType
}

func (p *PutStatement) Execute(memory *map[string]int) error {
	value, err := p.Expression.Evaluate(memory)
	if err != nil {
		return err
	}

	fmt.Println(value)
	return nil
}

// LetStatement
type LetStatement struct {
	Token tokens.Token
	Expression ExpressionType
}

func (s *LetStatement) Execute(memory *map[string]int) error {
	if memory == nil {
		return ErrMemoryNotInitialized
	}

	if _, exists := (*memory)[s.Token.Value]; exists {
		return ErrRedeclearationOfVariable
	}

	value, err := s.Expression.Evaluate(memory)
	if err != nil {
		return err
	}

	(*memory)[s.Token.Value] = value

	return nil
}

// IfStatement
type IfStatement struct {
	Condition ExpressionType
	Scope Scope
}

func (s *IfStatement) Execute(memory *map[string]int) error {
	condition, err := s.Condition.Evaluate(memory)

	if err != nil {
		return err
	}

	if condition == 0 {
		return nil
	} 

	for _, statement := range s.Scope.Statements {
		err := statement.Execute(memory)
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

func (s *ExitStatement) Execute(memory *map[string]int) error {
	exitCode, err := s.ExitCode.Evaluate(memory)
	if err != nil {
		return err
	}

	return ErrProgramExits(exitCode)
}
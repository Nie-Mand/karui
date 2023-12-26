package executor

import (
	"Nie-Mand/karui/core/parser"
)

type Executor struct {
	Program *parser.Program
}

func NewExecutor(_program *parser.Program) *Executor {
	return &Executor{
		Program: _program,
	}
}

func (e *Executor) ExecuteProgram() error {
	for _, statement := range e.Program.Statements {
		err := e.ExecuteStatement(statement)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *Executor) ExecuteStatement(statement parser.StatementType) error {
	return statement.Execute()
}

func (e *Executor) ExecuteScope() {}
func (e *Executor) ExecuteExpression() {}
func (e *Executor) ExecuteBinaryExpression() {}


func (e *Executor) ExecuteTerm(term parser.TermType) {

}
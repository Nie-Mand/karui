package executor

import (
	"github.com/Nie-Mand/karui/core/parser"
)

type Executor struct {
	Program *parser.Program
	Memory map[string]int
}

func NewExecutor(_program *parser.Program) *Executor {
	return &Executor{
		Program: _program,
		Memory: make(map[string]int, 0),
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
	return statement.Execute(&e.Memory)
}
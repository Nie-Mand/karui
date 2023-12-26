package parser

import "fmt"

type Program struct {
	Statements []StatementType
}

func NewProgram() *Program {
	return &Program{
		Statements: make([]StatementType, 0),
	}
}

func (p *Program) String() string {
	representation := "Program: \n"
	for idx, statement := range p.Statements {
		representation += fmt.Sprintf("%d) %s\n", idx, statement.String())
	}
	
	return representation
}
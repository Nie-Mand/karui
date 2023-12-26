package parser

import "fmt"

type Scope struct {
	Statements []StatementType
}

func NewScope() *Scope {
	return &Scope{
		Statements: []StatementType{},
	}
}

func (Scope) IsStatement() {}

func (s *Scope) String() string {
	representation := "Scope: "
	for idx, statement := range s.Statements {
		representation += fmt.Sprintf("%d) %s", idx, statement.String())
	}
	
	return representation
}
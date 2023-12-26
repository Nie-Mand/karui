package parser

type Program struct {
	Statements []StatementType
}

func NewProgram() *Program {
	return &Program{
		Statements: make([]StatementType, 0),
	}
}
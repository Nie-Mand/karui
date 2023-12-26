package parser

type Scope struct {
	Statements []StatementType
}

func NewScope() *Scope {
	return &Scope{
		Statements: []StatementType{},
	}
}

func (s *Scope) Execute(memory *map[string]int) error {
	for _, statement := range s.Statements {
		err := statement.Execute(memory)
		if err != nil {
			return err
		}
	}

	return nil
}
package parser

type Scope struct {
	Statements []StatementType
}

func NewScope() *Scope {
	return &Scope{
		Statements: []StatementType{},
	}
}

func (s *Scope) Execute() error {
	for _, statement := range s.Statements {
		err := statement.Execute()
		if err != nil {
			return err
		}
	}

	return nil
}
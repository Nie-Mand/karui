package tokens

type TokenType int

const (
	// Statements
	Exit TokenType = iota
	Let
	Puts
	If

	// Identifiers
	Identifier

	// Literals
	IntLiteral
	
	// Symbols
	OpenParenthesis
	CloseParenthesis
	OpenCurly
	CloseCurly
	Semicolon
	
	//
	Equal
	Plus
	Minus
	Multiply
	Divide
	Modulo
)

func (t TokenType) String() string {
	switch t {
	case Exit:
		return "exit"
	case Let:
		return "let"
	case Puts:
		return "puts"
	case If:
		return "if"

	case Identifier:
		return "identifier"
	case IntLiteral:
		return "int-literal"

	case Equal:
		return "="
	case Plus:
		return "+"
	case Minus:
		return "-"
	case Multiply:
		return "*"
	case Divide:
		return "/"
	case Modulo:
		return "%"
	case OpenParenthesis:
		return "("
	case CloseParenthesis:
		return ")"
	case OpenCurly:
		return "{"
	case CloseCurly:
		return "}"
	case Semicolon:
		return ";"
		
	default:
		return "Unknown"
	}
}